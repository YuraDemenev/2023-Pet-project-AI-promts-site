ENGLISH
# 2023-Pet-project-AI-promts-site
on my site you can upload images, click on images and view the promts you used to create them, or like the images you like. If you have registered you can view your profile where you can see all the images you have uploaded and liked in 2 different columns and additional information. You can use search for find image with special promt.

On the site you can upload or view images as an unregistered user, but you can't like them or access your profile.
Interesting implementation points:
- The site was loaded using docker compose and docker containers
- Password hashing
- infinite scroll down
- use of Middleware 
- use of HTMX to change elements on the site without reloading the page
- Image optimization (low quality images are loaded on the homepage for optimization, when the user clicks on the image the original image is loaded.
- Lazy loading when loading images (if traffic is slow a 20px blurry image is loaded)
- GIN framework 
- database postgres 
- Bootstrap for CSS
- Redis for caching images
- VPS server on linux ubuntu
- Nginx and SSL for site work

Site link: https://imagepromts.ru/

To raise the site locally you need to move the following files to an empty folder: docker-compose.yml, images folder, config folder, initdb folder, then write docker-compose up in the console.

РУССКИЙ
# 2023-Pet-project-AI-promts-site
На моем сайте вы можете загружать изображения, нажимать на них и просматривать промты, которые вы использовали для их создания, или ставить лайки понравившимся вам изображениям. Если вы зарегистрировались, вы можете просмотреть свой профиль, где вы можете увидеть все загруженные и понравившиеся вам изображения в двух разных колонках и дополнительную информацию. Вы можете воспользоваться поиском, чтобы найти изображение со специальным промтом.

На сайте вы можете загружать и просматривать изображения как незарегистрированный пользователь, но не можете ставить лайки или заходить в свой профиль.
Интересные моменты реализации:
- Сайт был загружен с помощью docker compose и docker containers
- хеширование пароля
- бесконечная прокрутка вниз
- использование Middleware 
- использование HTMX для изменения элементов на сайте без перезагрузки страницы
- Оптимизация изображений (изображения низкого качества загружаются на главную страницу для оптимизации, когда пользователь нажимает на изображение, загружается оригинальное изображение.
- Ленивая загрузка при загрузке изображений (при медленном трафике загружается 20px размытое изображение)
- фреймворк GIN 
- база данных postgres 
- Bootstrap для CSS
- Redis для кэширования изображений
- VPS сервер на linux ubuntu
- Nginx и SSL для работы сайта

Site link: https://imagepromts.ru/

Чтобы поднять сайт локально нужно переместить в пустую папку следующие файлы:docker-compose.yml, папку images,папку config,паку initdb, затем прописать в консоли docker-compose up
