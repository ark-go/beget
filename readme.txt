требуется доступ к DNS API Beget
требуется поддомен _acme-challenge.xxxx.ru


Удалите certbot-auto и любые пакеты ОС Certbot
Если  есть какие-либо пакеты Certbot, установленные с помощью диспетчера пакетов ОС, 
такого как apt, dnf или yum, вам следует удалить их перед установкой оснастки Certbot,
>apt-get remove certbot

>apt install snapd  -- Если не было
>snap install --classic certbot  -- Установка
при первом запуске просит указать почту,  если не было регистрации

Подготовьте команду Certbot
>ln -s /snap/bin/certbot /usr/bin/certbot

Для регистрации выполните команду:  
>certbot register -m admin@example.com  -- свой емайл

Запустите эту команду в командной строке на компьютере, чтобы подтвердить, 
что установленный подключаемый модуль будет иметь то же classic содержание, что и оснастка Certbot

--dry-run использовать флаг для проверкии тестирования , 
    без теста, при частых обращениях, Certbot обидится и заблокирует доступ
--cert-name  будет определять название катлога для хранения сертификатов, если его не указать, про обновлениях возможна смена каталога на xx001,xx002 и т.п.
--preferred-challenges=dns  метод получения сертификата
--manual   режим, говорит о том что надо запустить следущие хуки
--manual-auth-hook  указывается что запустить для записи ключей проверки certbot в DNS запись TXT (запускаем наш beget)
--manual-cleanup-hook  указывается программа которая сотрет записи TXT
--manual-public-ip-logging-ok для команды certbot при выполнении  --manual
    дает Let's Encrypt разрешение регистрировать общедоступный IP-адрес? не разбирался не проверял
-d  имена доменов 

проверяем:
root@zzzz:~# certbot certonly --dry-run --cert-name xxxxx.ru --manual --preferred-challenges=dns --manual-public-ip-logging-ok --manual-auth-hook "/root/beget/beget -save" --manual-cleanup-hook "/root/beget/beget -clear" -d xxxxx.ru,*.xxxxx.ru


и потом без флага

root@zzzz:~# certbot certonly --cert-name xxxxx.ru --manual --preferred-challenges=dns --manual-public-ip-logging-ok --manual-auth-hook "/root/beget/beget -save" --manual-cleanup-hook "/root/beget/beget -clear" -d xxxxx.ru,*.xxxxx.ru


файл beget.cfg рядом с исполняемым файлом:
UserLoginDns=xxxxxxx   # // Login для доступа к DNS
UserPasswdDns="xxxxxx" # // password для доступа к DNS
TimePropagation=120    # // таймаут в сек. для  DNS propagation  время для обновления записей ТХТ 

смотрим, проверяем права к файлам, Сertbot по умолчанию от root запускается


просто информация, данные от Certbot передаются через системные переменные
CERTBOT_DOMAIN
CERTBOT_VALIDATION