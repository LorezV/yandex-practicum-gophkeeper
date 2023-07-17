package usecase

import (
	"log"
	"os"

	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/LorezV/gophkeeper/internal/utils"
	"github.com/fatih/color"
	"github.com/google/uuid"
)

func (uc *GophKeeperClientUseCase) AddBinary(userPassword string, binary *entity.Binary) {
	accessToken, err := uc.authorisationCheck(userPassword)
	if err != nil {
		log.Fatalf("GophKeeperClientUseCase - authorisationCheck - %v", err)
	}
	file, err := os.Stat(binary.FileName)
	if err != nil {
		log.Fatalf("GophKeeperClientUseCase - AddBinary - %v", err)
	}
	tmpFilePath := "/tmp/" + file.Name()

	if err := utils.EncryptFile(userPassword, binary.FileName, tmpFilePath); err != nil {
		log.Fatalf("GophKeeperClientUseCase - EncryptFile - %v", err)
	}

	if err := uc.clientAPI.AddBinary(accessToken, binary, tmpFilePath); err != nil {
		log.Fatalf("GophKeeperClientUseCase - clientAPI.AddBinary - %v", err)
	}

	if err := uc.repo.AddBinary(binary); err != nil {
		log.Fatalf("GophKeeperClientUseCase - repo.AddBinary - %v", err)
	}

	if err := os.Rename(
		tmpFilePath,
		uc.cfg.FilesStorage.Location+"/"+binary.ID.String()); err != nil {
		log.Fatalf("GophKeeperClientUseCase - os.Rename - %v", err)
	}
	color.Green("Binary %v - %s saved", binary.ID, binary.FileName)
}

func (uc *GophKeeperClientUseCase) DelBinary(userPassword, binaryID string) {
	accessToken, err := uc.authorisationCheck(userPassword)
	if err != nil {
		return
	}
	binaryUUID, err := uuid.Parse(binaryID)
	if err != nil {
		color.Red(err.Error())
		log.Fatalf("GophKeeperClientUseCase - DelBinary - uuid.Parse - %v", err)
	}

	if err := uc.repo.DelBinary(binaryUUID); err != nil {
		log.Fatalf("GophKeeperClientUseCase - repo.DelNote - %v", err)
	}

	if err := uc.clientAPI.DelBinary(accessToken, binaryID); err != nil {
		log.Fatalf("GophKeeperClientUseCase - clientAPI.DelBinary - %v", err)
	}
	if err := os.Remove(uc.cfg.FilesStorage.Location + "/" + binaryID); err != nil {
		log.Fatalf("GophKeeperClientUseCase -  os.Remove - %v", err)
	}

	color.Green("Binary %q removed", binaryID)
}

func (uc *GophKeeperClientUseCase) GetBinary(userPassword, binaryID, filePath string) {
	accessToken, err := uc.authorisationCheck(userPassword)
	if err != nil {
		return
	}

	binaryUUID, err := uuid.Parse(binaryID)
	if err != nil {
		log.Fatalf("GophKeeperClientUseCase - GetBinary - uuid.Parse - %v", err)
	}

	binary, err := uc.repo.GetBinaryByID(binaryUUID)
	if err != nil {
		log.Fatalf("GophKeeperClientUseCase - GetBinary - GetBinaryByID - %v", err)
	}

	tmpFilePath := "/tmp/" + binary.FileName
	if err = uc.clientAPI.DownloadBinary(accessToken, tmpFilePath, &binary); err != nil {
		log.Fatalf("GophKeeperClientUseCase - GetBinary - clientAPI.DownloadBinary - %v", err)
	}

	if err = utils.DecryptFile(userPassword, tmpFilePath, filePath); err != nil {
		log.Fatalf("GophKeeperClientUseCase - GetBinary - EncryptFile - %v", err)
	}

	color.Green("File decrypted to %s", filePath)
	color.Green("%v", binary)
}

func (uc *GophKeeperClientUseCase) loadBinaries(accessToken string) {
	binaries, err := uc.clientAPI.GetBinaries(accessToken)
	if err != nil {
		color.Red("Connection error: %v", err)

		return
	}

	if err = uc.repo.SaveBinaries(binaries); err != nil {
		log.Println(err)

		return
	}
	color.Green("Loaded %v binaries", len(binaries))
}
