package secretm

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/vcuadrosv/gambituser/awsgo"
	"github.com/vcuadrosv/gambituser/models"
)

func GetSecret(nombreSecret string) (models.SecretRDSJSON, error) {
	var datosSecret models.SecretRDSJSON
	fmt.Println("> Pido secreto" + nombreSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})
	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println("> Lectura secret OK" + nombreSecret)
	return datosSecret, nil
}
