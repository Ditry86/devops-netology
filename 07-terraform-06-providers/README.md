# Домашнее задание к занятию "7.6. Написание собственных провайдеров для Terraform."

Бывает, что 
* общедоступная документация по терраформ ресурсам не всегда достоверна,
* в документации не хватает каких-нибудь правил валидации или неточно описаны параметры,
* понадобиться использовать провайдер без официальной документации,
* может возникнуть необходимость написать свой провайдер для системы используемой в ваших проектах.   

## Задача 1. 
Давайте потренируемся читать исходный код AWS провайдера, который можно склонировать от сюда: 
[https://github.com/hashicorp/terraform-provider-aws.git](https://github.com/hashicorp/terraform-provider-aws.git).
Просто найдите нужные ресурсы в исходном коде и ответы на вопросы станут понятны.  


1. Найдите, где перечислены все доступные `resource` и `data_source`, приложите ссылку на эти строки в коде на 
гитхабе.   
1. Для создания очереди сообщений SQS используется ресурс `aws_sqs_queue` у которого есть параметр `name`. 
    * С каким другим параметром конфликтует `name`? Приложите строчку кода, в которой это указано.
    * Какая максимальная длина имени? 
    * Какому регулярному выражению должно подчиняться имя? 

### **Ответ:**

1. Ресурсы AWS-провайдера:

    а) [Ссылка](https://github.com/hashicorp/terraform-provider-aws/blob/5b1b7fe382df81827632a72aca8bc7879a2957c5/internal/provider/provider.go#L426) на все `data_source`;

    б) [Ссылка](https://github.com/hashicorp/terraform-provider-aws/blob/5b1b7fe382df81827632a72aca8bc7879a2957c5/internal/provider/provider.go#L956) на все `resource`.

2. Ресурс `aws_sqs_queue`:

а) `name` конфликтует с параметром `name_prefix` (ConflictsWith: []string{"name_prefix"}). Фрагмент кода пакета sqs (файл queue.go):
        
```
    var (
	queueSchema = map[string]*schema.Schema{
    ...
        "name": {
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ForceNew:      true,
			ConflictsWith: []string{"name_prefix"},
		},
		"name_prefix": {
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ForceNew:      true,
			ConflictsWith: []string{"name"},
        },
    ...
    }
```


б) Максимальное кол-во символов (длина) для параметра `name` составляет 80. Данное ограничение указано в описании поля `QueueName` структуры `CreateQueueInput`, которая в свою очередь учавствует при создании resourse SQS.

Комментарии к коду [(https://github.com/aws/aws-sdk-go/blob/bda57440af7a0c67582025a7a5a5642de7cc9d6a/service/sqs/api.go#L2775)](https://github.com/aws/aws-sdk-go/blob/bda57440af7a0c67582025a7a5a5642de7cc9d6a/service/sqs/api.go#L2775):

```
// The name of the new queue. The following limits apply to this name:
	//
	//    * A queue name can have up to 80 characters.
	//
	//    * Valid values: alphanumeric characters, hyphens (-), and underscores
	//    (_).
	//
	//    * A FIFO queue name must end with the .fifo suffix.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueName is a required field
	QueueName *string `type:"string" required:"true"`
```

3) В исходниках провайдера и API AWS не нашел не одной функции, либо выражения, фильтрующих вводимое имя для SQS очереди, кроме функции проверяющей имя созданной очереди на содержание в нем уникального ID и суффикса [(https://github.com/hashicorp/terraform-provider-aws/blob/5800b3a7431e7f74ba40ec00411be14d27757372/internal/create/naming.go#L30)](https://github.com/hashicorp/terraform-provider-aws/blob/5800b3a7431e7f74ba40ec00411be14d27757372/internal/create/naming.go#L30). Но в описании того же `CreateQueueInput` указаны ограничения по вводимым символам:

```
//
	//    * Valid values: alphanumeric characters, hyphens (-), and underscores
	//    (_).
```
На основании этого можно сформулировать regexp, например, для маски вводимого значения имени очереди при его создании - `^[A-Za-z0-9_-]{1,80}$`.

## Задача 2. (Не обязательно) 
В рамках вебинара и презентации мы разобрали как создать свой собственный провайдер на примере кофемашины. 
Также вот официальная документация о создании провайдера: 
[https://learn.hashicorp.com/collections/terraform/providers](https://learn.hashicorp.com/collections/terraform/providers).

1. Проделайте все шаги создания провайдера.
2. В виде результата приложение ссылку на исходный код.
3. Попробуйте скомпилировать провайдер, если получится то приложите снимок экрана с командой и результатом компиляции.   

### **Не выполнено. Не успеваю до 25.12.2022 (закрытие модуля)**
---

### Как cдавать задание

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---
