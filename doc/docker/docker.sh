sudo docker run -d --restart always -p 3306:3306 --name mysql5.7 -v /var/lib/mysql:/var/lib/mysql -e MYSQL_ROOT_HOST=% -e MYSQL_ROOT_PASSWORD=root mysql:5.7
sudo docker run -d --restart always --name redis -p 6379:6379 redis --requirepass "pass"