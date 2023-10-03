# 2023-Pet-project-AI-promts-site
on my site you can upload images, click on images and view the promts you used to create them, or like the images you like. If you have registered you can view your profile where you can see all the images you have uploaded and liked in 2 different columns and additional information.
 On the site you can upload or view images as an unregistered user, but you can't like them or access your profile.
Interesting implementation points:
- Password hashing
- use of Middleware
- use of HTMX to change elements on the site without reloading the page
- Image optimization (low quality images are loaded on the homepage for optimization, when the user clicks on the image the original image is loaded.
- Lazy loading when loading images (if traffic is slow a 20px blurry image is loaded)
- GIN framework 
- database postgres 
- Bootstrap for CSS 

Open site:
If you want to open a site on your computer you need PostgreSQL to create a database there is init.up.sql in schema/ it will create the same database.!!! NOTE FOR THE SITE TO WORK YOU NEED ANONYM USER WITH ID=1 !!!. you can change the id of anonym in pkg\handler\middleware.go it is const you can change it.
