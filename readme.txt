

Удалите certbot-auto и любые пакеты ОС Certbot
Если  есть какие-либо пакеты Certbot, установленные с помощью диспетчера пакетов ОС, 
такого как apt, dnf или yum, вам следует удалить их перед установкой оснастки Certbot,

apt-get remove certbot

apt install snapd  -- Если не было
snap install --classic certbot  -- Установка
при первом запуске просит указать почту,  если не было регистрации

Подготовьте команду Certbot
ln -s /snap/bin/certbot /usr/bin/certbot

Для регистрации выполните команду:  
  certbot register -m admin@example.com  -- свой емайл

Запустите эту команду в командной строке на компьютере, чтобы подтвердить, 
что установленный подключаемый модуль будет иметь то же classic содержание, что и оснастка Certbot

для проверкии тестирования  использовать флаг --dry-run, 
без теста, при частых обращениях, Certbot обидится и заблокирует доступ

проверяем:
root@zzzz:~# certbot certonly --dry-run --cert-name xxxxx.ru --manual --preferred-challenges=dns --manual-public-ip-logging-ok --manual-auth-hook "/root/beget/beget -save" --manual-cleanup-hook "/root/beget/beget -clear" -d xxxxx.ru,*.xxxxx.ru

и потом без флага

root@zzzz:~# certbot certonly --cert-name xxxxx.ru --manual --preferred-challenges=dns --manual-public-ip-logging-ok --manual-auth-hook "/root/beget/beget -save" --manual-cleanup-hook "/root/beget/beget -clear" -d xxxxx.ru,*.xxxxx.ru

