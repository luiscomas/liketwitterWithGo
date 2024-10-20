package secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/luiscomas/liketwitterWithGo/awsgo"
	"github.com/luiscomas/liketwitterWithGo/models"
)

// GetSecrets retrieves the secret from AWS Secrets Manager
func GetSecret(SecretName string) (models.Secret, error) {
	var datosSecret models.Secret
	fmt.Println("SecretName: ", SecretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(SecretName),
	})
	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println("Lectura de Secret OK " + SecretName)
	return datosSecret, nil
}
