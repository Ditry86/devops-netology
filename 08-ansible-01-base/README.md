# Домашнее задание к занятию "1. Введение в Ansible"

## Подготовка к выполнению
1. Установите ansible версии 2.10 или выше.
2. Создайте свой собственный публичный репозиторий на github с произвольным именем.
3. Скачайте [playbook](./playbook/) из репозитория с домашним заданием и перенесите его в свой репозиторий.

## Основная часть
1. Попробуйте запустить playbook на окружении из `test.yml`, зафиксируйте какое значение имеет факт `some_fact` для указанного хоста при выполнении playbook'a.

### **Ответ:**

```
TASK [Print fact] **********************************************************************************************************************************************************************************************************************************************************
ok: [localhost] => {
    "msg": 12
}
```

2. Найдите файл с переменными (group_vars) в котором задаётся найденное в первом пункте значение и поменяйте его на 'all default fact'.
### **Ответ:**

group_vars/all/examp.yml:

```yml
  some_fact: 'all default fact'
```

3. Воспользуйтесь подготовленным (используется `docker`) или создайте собственное окружение для проведения дальнейших испытаний.

### **Ответ:**

Для окружения docker создал два контейнера на локальном хосте с помощью docer-compose. Поскольку в оффициальном образе ubuntu отсутствует библиотеки и интерпертатор Python, для работы ansible необходимо установить их. Несколько способов сделать это (но не все): использовать другой образ Ubuntu с предустановленным Python, используя Ansible raw module в playbook установить Python на запущенном контейнере, используя Dockerfile добавить слой с утсановкой пакета Python. Я выбрал последний.
Подготовленное окружение docker:

    1) docker/compose.yml (docker-compose file):

    ```yml
    services:

        centos7:
            image: centos
            container_name: centos7
            networks: 
            my_net:
                ipv4_address: 192.168.2.10
            tty: true

        ubuntu:
            build: 
            context: .
            dockerfile: Dockerfile
            container_name: ubuntu
            networks: 
            my_net:
                ipv4_address: 192.168.2.11
            tty: true

    networks:

        my_net:
            driver: bridge
            ipam:
            config:
            - subnet: 192.168.2.0/24
                gateway: 192.168.2.254
    ```
    2) docker/Dockerfile:
    
    ```
    FROM ubuntu
    RUN apt update && apt install -y --force-yes python3
    ```

4. Проведите запуск playbook на окружении из `prod.yml`. Зафиксируйте полученные значения `some_fact` для каждого из `managed host`.

### **Ответ:**

```
...
PLAY RECAP *****************************************************************************************************************************************************************************************************************************************************************
centos7                    : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   
ubuntu                     : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0 
```

5. Добавьте факты в `group_vars` каждой из групп хостов так, чтобы для `some_fact` получились следующие значения: для `deb` - 'deb default fact', для `el` - 'el default fact'.
6.  Повторите запуск playbook на окружении `prod.yml`. Убедитесь, что выдаются корректные значения для всех хостов.

### **Ответ:**

```
...
TASK [Print fact] **********************************************************************************************************************************************************************************************************************************************************
ok: [centos7] => {
    "msg": "el default fact"
}
ok: [ubuntu] => {
    "msg": "deb default fact"
}

PLAY RECAP *****************************************************************************************************************************************************************************************************************************************************************
centos7                    : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   
ubuntu                     : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0  
```

7. При помощи `ansible-vault` зашифруйте факты в `group_vars/deb` и `group_vars/el` с паролем `netology`.

### **Ответ:**

```bash
ansible-vault encrypt group_vars/deb/examp.yml
...

ansible-vault encrypt group_vars/el/examp.yml
...

cat group_vars/deb/examp.yml
$ANSIBLE_VAULT;1.1;AES256
36353637363334643537616336336366666138303361323134646431653634356130346137333865
3336353535633330653239323165393161646163666363650a643165633332323334616263633631
62316633363839383336363736323930363236323436613839613232336363313030393939373234
6633303033643933650a616464636561373462303435333137373664356130323963373335653530
64353237306531343062623231343336633736363637623934313535323630343362303730363038
6266376264346263323035333535633263343664363332393038

cat group_vars/el/examp.yml
$ANSIBLE_VAULT;1.1;AES256
62373434623135613134396537663938663061646566373039636634616439373636626461363263
3665383866643566363666613738646338343532613064320a303737656239666563306664326534
34373038643436363065666230313361383365313534343262666161613233326530643261343634
3835303237353261630a386231313131623237373636313562633464353465396538623132366262
62373737303962643763393364326663333339363232346133336265363039353034336466643362
3261663537373331376465656230663138613463366535363066
```

8. Запустите playbook на окружении `prod.yml`. При запуске `ansible` должен запросить у вас пароль. Убедитесь в работоспособности.

### **Ответ:**

```
ansible-playbook site.yml -i inventory/prod.yml --ask-vault-password
```
9. Посмотрите при помощи `ansible-doc` список плагинов для подключения. Выберите подходящий для работы на `control node`.

### **Ответ:**

Для работы с локальным хостом, который является `control node` можно использовать предустановленный `connection plugin` - `ansible.builtin.local` (либо просто `local`). 

10. В `prod.yml` добавьте новую группу хостов с именем  `local`, в ней разместите localhost с необходимым типом подключения.

### **Ответ:**

inventory/prod.yml:

```yml
---
  el:
    hosts:
      centos7:
        ansible_connection: docker
        
  deb:
    hosts:
      ubuntu:
        ansible_connection: docker

  local:
    hosts:
      localhost:
        ansible_connection: local
```

11. Запустите playbook на окружении `prod.yml`. При запуске `ansible` должен запросить у вас пароль. Убедитесь что факты `some_fact` для каждого из хостов определены из верных `group_vars`.

### **Ответ:**

```
ansible-playbook site.yml -i inventory/prod.yml --ask-vault-password
...
TASK [Print fact] **********************************************************************************************************************************************************************************************************************************************************
ok: [centos7] => {
    "msg": "el default fact"
}
ok: [ubuntu] => {
    "msg": "deb default fact"
}
ok: [localhost] => {
    "msg": "all default fact"
}

PLAY RECAP *****************************************************************************************************************************************************************************************************************************************************************
centos7                    : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   
localhost                  : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   
ubuntu                     : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   
```

12. Заполните `README.md` ответами на вопросы. Сделайте `git push` в ветку `master`. В ответе отправьте ссылку на ваш открытый репозиторий с изменённым `playbook` и заполненным `README.md`.

## **Ответ:**

[My Ansible repo (master branch)](https://github.com/Ditry86/ansible_study/tree/master)


## Необязательная часть

1. При помощи `ansible-vault` расшифруйте все зашифрованные файлы с переменными.
2. Зашифруйте отдельное значение `PaSSw0rd` для переменной `some_fact` паролем `netology`. Добавьте полученное значение в `group_vars/all/exmp.yml`.
3. Запустите `playbook`, убедитесь, что для нужных хостов применился новый `fact`.
4. Добавьте новую группу хостов `fedora`, самостоятельно придумайте для неё переменную. В качестве образа можно использовать [этот](https://hub.docker.com/r/pycontribs/fedora).
5. Напишите скрипт на bash: автоматизируйте поднятие необходимых контейнеров, запуск ansible-playbook и остановку контейнеров.
6. Все изменения должны быть зафиксированы и отправлены в вашей личный репозиторий.

## **Ответ:**

[My Ansible repo](https://github.com/Ditry86/ansible_study)

---

### Как оформить ДЗ?

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---