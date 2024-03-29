PGDMP     6    6                |            SiteWatchImage    15.4    15.4 &    %           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            &           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            '           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            (           1262    16709    SiteWatchImage    DATABASE     �   CREATE DATABASE "SiteWatchImage" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_United States.1252';
     DROP DATABASE "SiteWatchImage";
                postgres    false            �            1259    18489    consideration    TABLE     v   CREATE TABLE public.consideration (
    id bigint NOT NULL,
    user_id bigint,
    image_url text,
    title text
);
 !   DROP TABLE public.consideration;
       public         heap    postgres    false            �            1259    18488    consideration_id_seq    SEQUENCE     }   CREATE SEQUENCE public.consideration_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 +   DROP SEQUENCE public.consideration_id_seq;
       public          postgres    false    221            )           0    0    consideration_id_seq    SEQUENCE OWNED BY     M   ALTER SEQUENCE public.consideration_id_seq OWNED BY public.consideration.id;
          public          postgres    false    220            �            1259    18444    images    TABLE     �   CREATE TABLE public.images (
    user_id bigint NOT NULL,
    image_id bigint NOT NULL,
    image_url text NOT NULL,
    like_count integer,
    title text
);
    DROP TABLE public.images;
       public         heap    postgres    false            �            1259    18443    images_image_id_seq    SEQUENCE     |   CREATE SEQUENCE public.images_image_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 *   DROP SEQUENCE public.images_image_id_seq;
       public          postgres    false    217            *           0    0    images_image_id_seq    SEQUENCE OWNED BY     K   ALTER SEQUENCE public.images_image_id_seq OWNED BY public.images.image_id;
          public          postgres    false    216            �            1259    18471    likes    TABLE     p   CREATE TABLE public.likes (
    user_id bigint NOT NULL,
    image_url text NOT NULL,
    like_check boolean
);
    DROP TABLE public.likes;
       public         heap    postgres    false            �            1259    18459    promts    TABLE     U   CREATE TABLE public.promts (
    title text NOT NULL,
    image_url text NOT NULL
);
    DROP TABLE public.promts;
       public         heap    postgres    false            �            1259    18154    users    TABLE     �   CREATE TABLE public.users (
    id bigint NOT NULL,
    username text NOT NULL,
    password_hash character varying(255) NOT NULL
);
    DROP TABLE public.users;
       public         heap    postgres    false            �            1259    18153    users_id_seq    SEQUENCE     u   CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.users_id_seq;
       public          postgres    false    215            +           0    0    users_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
          public          postgres    false    214            y           2604    18492    consideration id    DEFAULT     t   ALTER TABLE ONLY public.consideration ALTER COLUMN id SET DEFAULT nextval('public.consideration_id_seq'::regclass);
 ?   ALTER TABLE public.consideration ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    220    221    221            x           2604    18447    images image_id    DEFAULT     r   ALTER TABLE ONLY public.images ALTER COLUMN image_id SET DEFAULT nextval('public.images_image_id_seq'::regclass);
 >   ALTER TABLE public.images ALTER COLUMN image_id DROP DEFAULT;
       public          postgres    false    217    216    217            w           2604    18157    users id    DEFAULT     d   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    215    214    215            "          0    18489    consideration 
   TABLE DATA           F   COPY public.consideration (id, user_id, image_url, title) FROM stdin;
    public          postgres    false    221   �)                 0    18444    images 
   TABLE DATA           Q   COPY public.images (user_id, image_id, image_url, like_count, title) FROM stdin;
    public          postgres    false    217   *                  0    18471    likes 
   TABLE DATA           ?   COPY public.likes (user_id, image_url, like_check) FROM stdin;
    public          postgres    false    219   *                 0    18459    promts 
   TABLE DATA           2   COPY public.promts (title, image_url) FROM stdin;
    public          postgres    false    218   ;*                 0    18154    users 
   TABLE DATA           <   COPY public.users (id, username, password_hash) FROM stdin;
    public          postgres    false    215   X*       ,           0    0    consideration_id_seq    SEQUENCE SET     C   SELECT pg_catalog.setval('public.consideration_id_seq', 1, false);
          public          postgres    false    220            -           0    0    images_image_id_seq    SEQUENCE SET     B   SELECT pg_catalog.setval('public.images_image_id_seq', 1, false);
          public          postgres    false    216            .           0    0    users_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.users_id_seq', 4, true);
          public          postgres    false    214            �           2606    18496     consideration consideration_pkey 
   CONSTRAINT     ^   ALTER TABLE ONLY public.consideration
    ADD CONSTRAINT consideration_pkey PRIMARY KEY (id);
 J   ALTER TABLE ONLY public.consideration DROP CONSTRAINT consideration_pkey;
       public            postgres    false    221                       2606    18453    images images_image_url_key 
   CONSTRAINT     [   ALTER TABLE ONLY public.images
    ADD CONSTRAINT images_image_url_key UNIQUE (image_url);
 E   ALTER TABLE ONLY public.images DROP CONSTRAINT images_image_url_key;
       public            postgres    false    217            �           2606    18451    images images_pkey 
   CONSTRAINT     `   ALTER TABLE ONLY public.images
    ADD CONSTRAINT images_pkey PRIMARY KEY (user_id, image_url);
 <   ALTER TABLE ONLY public.images DROP CONSTRAINT images_pkey;
       public            postgres    false    217    217            �           2606    18477    likes likes_pkey 
   CONSTRAINT     ^   ALTER TABLE ONLY public.likes
    ADD CONSTRAINT likes_pkey PRIMARY KEY (user_id, image_url);
 :   ALTER TABLE ONLY public.likes DROP CONSTRAINT likes_pkey;
       public            postgres    false    219    219            �           2606    18465    promts promts_pkey 
   CONSTRAINT     ^   ALTER TABLE ONLY public.promts
    ADD CONSTRAINT promts_pkey PRIMARY KEY (title, image_url);
 <   ALTER TABLE ONLY public.promts DROP CONSTRAINT promts_pkey;
       public            postgres    false    218    218            {           2606    18161    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    215            }           2606    18163    users users_username_key 
   CONSTRAINT     W   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);
 B   ALTER TABLE ONLY public.users DROP CONSTRAINT users_username_key;
       public            postgres    false    215            �           2606    18497 (   consideration consideration_user_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.consideration
    ADD CONSTRAINT consideration_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;
 R   ALTER TABLE ONLY public.consideration DROP CONSTRAINT consideration_user_id_fkey;
       public          postgres    false    3195    221    215            �           2606    18454    images images_user_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.images
    ADD CONSTRAINT images_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;
 D   ALTER TABLE ONLY public.images DROP CONSTRAINT images_user_id_fkey;
       public          postgres    false    3195    215    217            �           2606    18483    likes likes_image_url_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.likes
    ADD CONSTRAINT likes_image_url_fkey FOREIGN KEY (image_url) REFERENCES public.images(image_url) ON DELETE CASCADE;
 D   ALTER TABLE ONLY public.likes DROP CONSTRAINT likes_image_url_fkey;
       public          postgres    false    217    219    3199            �           2606    18478    likes likes_user_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.likes
    ADD CONSTRAINT likes_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;
 B   ALTER TABLE ONLY public.likes DROP CONSTRAINT likes_user_id_fkey;
       public          postgres    false    215    219    3195            �           2606    18466    promts promts_image_url_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.promts
    ADD CONSTRAINT promts_image_url_fkey FOREIGN KEY (image_url) REFERENCES public.images(image_url) ON DELETE CASCADE;
 F   ALTER TABLE ONLY public.promts DROP CONSTRAINT promts_image_url_fkey;
       public          postgres    false    217    3199    218            "      x������ � �            x������ � �             x������ � �            x������ � �         �   x���1!F�z9�C�?	����X��x{w����m�������A�5#[���$\�.B�1	3�Õ(h�|<��1�5����H�]�q�O�Zk��<�X��U&�����lN����}ʖ�Q�g�:����iR\ZO����vI)}��\�     