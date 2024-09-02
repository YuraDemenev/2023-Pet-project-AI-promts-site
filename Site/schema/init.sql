--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4
-- Dumped by pg_dump version 16.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: consideration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.consideration (
    id bigint NOT NULL,
    user_id bigint,
    image_url text,
    title text
);


ALTER TABLE public.consideration OWNER TO postgres;

--
-- Name: consideration_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.consideration_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.consideration_id_seq OWNER TO postgres;

--
-- Name: consideration_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.consideration_id_seq OWNED BY public.consideration.id;


--
-- Name: images; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.images (
    user_id bigint NOT NULL,
    image_id bigint NOT NULL,
    image_url text NOT NULL,
    like_count integer,
    title text
);


ALTER TABLE public.images OWNER TO postgres;

--
-- Name: images_image_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.images_image_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.images_image_id_seq OWNER TO postgres;

--
-- Name: images_image_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.images_image_id_seq OWNED BY public.images.image_id;


--
-- Name: likes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.likes (
    user_id bigint NOT NULL,
    image_url text NOT NULL,
    like_check boolean
);


ALTER TABLE public.likes OWNER TO postgres;

--
-- Name: promts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.promts (
    title text NOT NULL,
    image_url text NOT NULL
);


ALTER TABLE public.promts OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    username text NOT NULL,
    password_hash character varying(255) NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: consideration id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.consideration ALTER COLUMN id SET DEFAULT nextval('public.consideration_id_seq'::regclass);


--
-- Name: images image_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.images ALTER COLUMN image_id SET DEFAULT nextval('public.images_image_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: consideration; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.consideration (id, user_id, image_url, title) FROM stdin;
\.


--
-- Data for Name: images; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.images (user_id, image_id, image_url, like_count, title) FROM stdin;
3	6	202408241601052417.jpg	0	a beautiful painting of a singular little cute and charming hedgehog with a red santa klaus hat in a snowy field, art by greg rutkowski and thomas kinkade, trending on artstation
3	7	20240824160113310.jpg	0	ciudad blanca marmol con esculturas de marmol griegas, de noche, luces negras, concept art, 2d
3	8	202408241601294801.jpg	0	an ultrarealistic photograph of an enormous minimalist futuristic full-body power suit, without lights, made of white metal and polymer, full crystal hull, for a muscular man, imposing, modern minimalism, elegant, dystopia, mysterious, godlike, scary, in a dark room, with a red cape, 8k
3	9	202408241601406977.jpg	0	pen and ink, illustrated by hergé, of a cat sleeping on a porch with lots of plants. stunning color scheme, masterpiece
3	10	202408241601491083.jpg	0	detailed, vibrant illustration of niagara falls, vivid and vibrant colors, well-detailed, summer, sunny day, realistic and beautiful illustration, style matisse.
3	1	202408241559475981.jpg	1	a neon-lit jungle with glowing flora and fauna, where the trees are circuit boards and the rivers flow with liquid light
3	3	202408241600179715.jpg	1	black background with smoke in spotlight, vibrant, colorful gradient splash, hd, 4k, high-quality, highly detailed, photorealistic, raw, high quality, dynamic lighting, sharp focus, ultra realistic, masterpiece
3	2	202408241600027195.jpg	1	trippy anime wonderland character  black background dominant black and blue colors halloweens cheshire cat
3	5	202408241600572009.jpg	1	zaha hadid  expressionist exterior beach by the ocean made of wood and concrate.
3	4	202408241600312513.jpg	1	beautiful san francisco, illustrated by hergé, style of tin tin comics, pen and ink., beautiful colors, attention to detail
4	11	202408241625226672.jpg	0	national geographic award winning drone photograph of a humpback whale spraying and spouting water above the surface, exciting movement, bright light, film grain, lens flare, bright morning sky, kodachrome iso 200
4	12	202408241625296638.jpg	0	extremely ultra-realistic photorealistic 3d, professional photography, natural lighting, volumetric lighting maximalist photo-illustration in 8k resolution, detailed, elegant an inverted flying island
4	13	20240824162537955.jpg	0	gotham city skyline at night, zbrush render, 8k
4	14	20240824162544402.jpg	0	an ultrarealistic photograph of a minimalist futuristic full body power suit, without lights, made of black mate metal and polymer, full crystal hull, woman shape, imposing, modern minimalism, elegant, dystopia, mysterious, godlike, scary, in a dark room, 8k
4	15	202408241625529530.jpg	0	the pixelated beach scene with a vibrant orange and pink sunset, silhouetted palm trees swaying in the breeze is a beautiful and serene sight.
4	16	202408241625591185.jpg	0	photo of concept car from 2025,8k resoultion,hyper realstic, black, night, 35mm film, editorial, high fashion, need for speed
4	17	202408241626056496.jpg	0	a cute grey cat in a boat floating in a starry night sky over a snowy city with a cute little boy wearing pajamas studio ghibli style. white background
4	18	20240824162614891.jpg	0	a knight on a snow covered mountain in front of a castle, in the style of goblincore, eroded interiors, intel core, dark cyan and orange, 32k uhd, made of wrought iron, heavy shading
4	19	20240824162630867.jpg	0	futuristic starship ,hyper realistic, detailed render, extremely complex and advanced chassis, natural dirt and debris detail, scuffs, , stunning details, dark hues, flying in the space
4	20	202408241626402517.jpg	0	a high-tech government center of a country located in the center of the country on a glacier with a floating island used to store important documents, holographic dome, fire station, police station, hospital, train station, and military base.
4	21	202408241626486334.jpg	0	a dragon on a black background with a sunrise in the background, a silk screen by sōami, reddit contest spring, synchromism, anime aesthetic, sky background, vaporwave
4	22	202408241626581896.jpg	0	a statue from a muscular greek god statue with a wavey middel parting hair out of white marbel with a black backround standing on a podest animeted style 8k
1	23	202408241627401778.jpg	0	vivid vision of colors that swirl into each other and create psychedelic forms, pastel, colors blue, light blue
1	24	202408241627497442.jpg	0	cube cutout of an isometric programmer's bedroom, 3d art, professional colors, soft lighting, high detail, artstation, concept art, behance, ray tracing
1	25	202408241627587662.jpg	0	all glassblowing vampire urban style shooting stars big crescent moon one white wolf one black wolf howling at the moon constellations night moon spark of blue light bats red purple blue cream
1	26	202408241628056788.jpg	0	wild west gunslinger, male, wild west outlaw, portrait, studio ghibli, akira toriyama, james gilleard, trending pixiv fanboxe, 4k, the style of 90's vintage anime, brown hair
1	27	20240824162826948.jpg	0	anime art of akira, detailed scene, red, perfect face, intricately detailed photorealism, trending on artstation, neon lights, rainy day, ray-traced environment, vintage 90's anime artwork.
1	28	202408241628342347.jpg	0	the joker walking through streets of new york, stunning photo, dark moody aesthetic, at night, city lights in background. surreal, 8k
1	29	202408241628447113.jpg	0	shabby chic, dreamy mist, pastel junk journals, christmas street with cafe and shops, swirling magical fairytale abstract art style.
1	30	202408241628518681.jpg	0	warhammer 40k chaos armor, black hair, jaw made from metal, holding a bolt gun
1	31	202408241628584219.jpg	0	a vertical internal circular cylindre corridor with minimalist space, an ogive circular roof ,white sand
1	32	202408241629068500.jpg	0	beautiful illustration of a man and woman in a pink dress watching the sunset setting over the golden gate bridge, golden hour, beautiful stunning color scheme, masterpiece
1	33	202408241629149280.jpg	0	a stunning iridescent male marble bust, very reflective, perfect lighting, dark background, zbrush
1	34	202408241629217321.jpg	0	a realistic painting of a stack of pancakes, bacon, eggs with complementary colors.
1	35	202408241629272060.jpg	0	beautiful cozy bedroom with floor to ceiling glass windows overlooking a cyberpunk city at night, thunderstorm outside with torrential rain, detailed, high resolution, photorrealistic, dark, gloomy, moody aesthetic
1	36	202408241629372780.jpg	0	create oil painting surrealistic in michael parkes style, raspberries, cherries on cheese cake with marmelade, background in pale blue and gold geometric pattern like vasily kandinsky composition 8, please no woman faces, only berries, stilllife in gustav klimt style, michael parkes style, john william waterhouse style, very high resolution, very realistic, 40 k
1	37	202408241629432651.jpg	0	detailed, vibrant illustration of a cowboy in the copper canyons the sierra of chihuahua state, by herge, in the style of tin-tin comics, vibrant colors, detailed, sunny day, attention to detail, 8k.
1	38	202408241629506434.jpg	0	a sphere shaped like a drought, with the lower part resembling water, a dripping drop falls from the sphere, set against a blue sky with white clouds, surrealism art
1	39	202408241630061375.jpg	0	a back vew of a mage man walking in a cobbeled street of a medieval village covered of snow during the night. he wearing a hooded coat.
1	40	2024082416301216.jpg	0	futuristic sci-fi pod chair, flat design, product-view, editorial photography, transparent orb, product photography, natural lighting, plants, natural daytime lighting, zbrush, 8k, natural wooden environment
\.


--
-- Data for Name: likes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.likes (user_id, image_url, like_check) FROM stdin;
4	202408241559475981.jpg	t
4	202408241600179715.jpg	t
4	202408241600027195.jpg	t
4	202408241600572009.jpg	t
4	202408241600312513.jpg	t
\.


--
-- Data for Name: promts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.promts (title, image_url) FROM stdin;
neon-lit	202408241559475981.jpg
jungle	202408241559475981.jpg
with	202408241559475981.jpg
glowing	202408241559475981.jpg
flora	202408241559475981.jpg
fauna	202408241559475981.jpg
where	202408241559475981.jpg
the	202408241559475981.jpg
trees	202408241559475981.jpg
are	202408241559475981.jpg
circuit	202408241559475981.jpg
boards	202408241559475981.jpg
rivers	202408241559475981.jpg
flow	202408241559475981.jpg
liquid	202408241559475981.jpg
light	202408241559475981.jpg
trippy	202408241600027195.jpg
anime	202408241600027195.jpg
wonderland	202408241600027195.jpg
character	202408241600027195.jpg
black	202408241600027195.jpg
background	202408241600027195.jpg
dominant	202408241600027195.jpg
blue	202408241600027195.jpg
colors	202408241600027195.jpg
halloweens	202408241600027195.jpg
cheshire	202408241600027195.jpg
cat	202408241600027195.jpg
black	202408241600179715.jpg
background	202408241600179715.jpg
with	202408241600179715.jpg
smoke	202408241600179715.jpg
in	202408241600179715.jpg
spotlight	202408241600179715.jpg
vibrant	202408241600179715.jpg
colorful	202408241600179715.jpg
gradient	202408241600179715.jpg
splash	202408241600179715.jpg
hd	202408241600179715.jpg
4k	202408241600179715.jpg
high-quality	202408241600179715.jpg
highly	202408241600179715.jpg
detailed	202408241600179715.jpg
photorealistic	202408241600179715.jpg
raw	202408241600179715.jpg
high	202408241600179715.jpg
quality	202408241600179715.jpg
dynamic	202408241600179715.jpg
lighting	202408241600179715.jpg
sharp	202408241600179715.jpg
focus	202408241600179715.jpg
ultra	202408241600179715.jpg
realistic	202408241600179715.jpg
masterpiece	202408241600179715.jpg
beautiful	202408241600312513.jpg
san	202408241600312513.jpg
francisco	202408241600312513.jpg
illustrated	202408241600312513.jpg
by	202408241600312513.jpg
herg	202408241600312513.jpg
style	202408241600312513.jpg
tin	202408241600312513.jpg
comics	202408241600312513.jpg
pen	202408241600312513.jpg
ink	202408241600312513.jpg
colors	202408241600312513.jpg
attention	202408241600312513.jpg
to	202408241600312513.jpg
detail	202408241600312513.jpg
zaha	202408241600572009.jpg
hadid	202408241600572009.jpg
expressionist	202408241600572009.jpg
exterior	202408241600572009.jpg
beach	202408241600572009.jpg
by	202408241600572009.jpg
the	202408241600572009.jpg
ocean	202408241600572009.jpg
made	202408241600572009.jpg
wood	202408241600572009.jpg
concrate	202408241600572009.jpg
beautiful	202408241601052417.jpg
painting	202408241601052417.jpg
singular	202408241601052417.jpg
little	202408241601052417.jpg
cute	202408241601052417.jpg
charming	202408241601052417.jpg
hedgehog	202408241601052417.jpg
with	202408241601052417.jpg
red	202408241601052417.jpg
santa	202408241601052417.jpg
klaus	202408241601052417.jpg
hat	202408241601052417.jpg
in	202408241601052417.jpg
snowy	202408241601052417.jpg
field	202408241601052417.jpg
art	202408241601052417.jpg
by	202408241601052417.jpg
greg	202408241601052417.jpg
rutkowski	202408241601052417.jpg
thomas	202408241601052417.jpg
kinkade	202408241601052417.jpg
trending	202408241601052417.jpg
on	202408241601052417.jpg
artstation	202408241601052417.jpg
ciudad	20240824160113310.jpg
blanca	20240824160113310.jpg
marmol	20240824160113310.jpg
con	20240824160113310.jpg
esculturas	20240824160113310.jpg
de	20240824160113310.jpg
griegas	20240824160113310.jpg
noche	20240824160113310.jpg
luces	20240824160113310.jpg
negras	20240824160113310.jpg
concept	20240824160113310.jpg
art	20240824160113310.jpg
2d	20240824160113310.jpg
an	202408241601294801.jpg
ultrarealistic	202408241601294801.jpg
photograph	202408241601294801.jpg
enormous	202408241601294801.jpg
minimalist	202408241601294801.jpg
futuristic	202408241601294801.jpg
full-body	202408241601294801.jpg
power	202408241601294801.jpg
suit	202408241601294801.jpg
without	202408241601294801.jpg
lights	202408241601294801.jpg
made	202408241601294801.jpg
white	202408241601294801.jpg
metal	202408241601294801.jpg
polymer	202408241601294801.jpg
full	202408241601294801.jpg
crystal	202408241601294801.jpg
hull	202408241601294801.jpg
for	202408241601294801.jpg
muscular	202408241601294801.jpg
man	202408241601294801.jpg
imposing	202408241601294801.jpg
modern	202408241601294801.jpg
minimalism	202408241601294801.jpg
elegant	202408241601294801.jpg
dystopia	202408241601294801.jpg
mysterious	202408241601294801.jpg
godlike	202408241601294801.jpg
scary	202408241601294801.jpg
in	202408241601294801.jpg
dark	202408241601294801.jpg
room	202408241601294801.jpg
with	202408241601294801.jpg
red	202408241601294801.jpg
cape	202408241601294801.jpg
8k	202408241601294801.jpg
pen	202408241601406977.jpg
ink	202408241601406977.jpg
illustrated	202408241601406977.jpg
by	202408241601406977.jpg
herg	202408241601406977.jpg
cat	202408241601406977.jpg
sleeping	202408241601406977.jpg
on	202408241601406977.jpg
porch	202408241601406977.jpg
with	202408241601406977.jpg
lots	202408241601406977.jpg
plants	202408241601406977.jpg
stunning	202408241601406977.jpg
color	202408241601406977.jpg
scheme	202408241601406977.jpg
masterpiece	202408241601406977.jpg
detailed	202408241601491083.jpg
vibrant	202408241601491083.jpg
illustration	202408241601491083.jpg
niagara	202408241601491083.jpg
falls	202408241601491083.jpg
vivid	202408241601491083.jpg
colors	202408241601491083.jpg
well-detailed	202408241601491083.jpg
summer	202408241601491083.jpg
sunny	202408241601491083.jpg
day	202408241601491083.jpg
realistic	202408241601491083.jpg
beautiful	202408241601491083.jpg
style	202408241601491083.jpg
matisse	202408241601491083.jpg
national	202408241625226672.jpg
geographic	202408241625226672.jpg
award	202408241625226672.jpg
winning	202408241625226672.jpg
drone	202408241625226672.jpg
photograph	202408241625226672.jpg
humpback	202408241625226672.jpg
whale	202408241625226672.jpg
spraying	202408241625226672.jpg
spouting	202408241625226672.jpg
water	202408241625226672.jpg
above	202408241625226672.jpg
the	202408241625226672.jpg
surface	202408241625226672.jpg
exciting	202408241625226672.jpg
movement	202408241625226672.jpg
bright	202408241625226672.jpg
light	202408241625226672.jpg
film	202408241625226672.jpg
grain	202408241625226672.jpg
lens	202408241625226672.jpg
flare	202408241625226672.jpg
morning	202408241625226672.jpg
sky	202408241625226672.jpg
kodachrome	202408241625226672.jpg
iso	202408241625226672.jpg
200	202408241625226672.jpg
extremely	202408241625296638.jpg
ultra-realistic	202408241625296638.jpg
photorealistic	202408241625296638.jpg
3d	202408241625296638.jpg
professional	202408241625296638.jpg
photography	202408241625296638.jpg
natural	202408241625296638.jpg
lighting	202408241625296638.jpg
volumetric	202408241625296638.jpg
maximalist	202408241625296638.jpg
photo-illustration	202408241625296638.jpg
in	202408241625296638.jpg
8k	202408241625296638.jpg
resolution	202408241625296638.jpg
detailed	202408241625296638.jpg
elegant	202408241625296638.jpg
an	202408241625296638.jpg
inverted	202408241625296638.jpg
flying	202408241625296638.jpg
island	202408241625296638.jpg
gotham	20240824162537955.jpg
city	20240824162537955.jpg
skyline	20240824162537955.jpg
at	20240824162537955.jpg
night	20240824162537955.jpg
zbrush	20240824162537955.jpg
render	20240824162537955.jpg
8k	20240824162537955.jpg
an	20240824162544402.jpg
ultrarealistic	20240824162544402.jpg
photograph	20240824162544402.jpg
minimalist	20240824162544402.jpg
futuristic	20240824162544402.jpg
full	20240824162544402.jpg
body	20240824162544402.jpg
power	20240824162544402.jpg
suit	20240824162544402.jpg
without	20240824162544402.jpg
lights	20240824162544402.jpg
made	20240824162544402.jpg
black	20240824162544402.jpg
mate	20240824162544402.jpg
metal	20240824162544402.jpg
polymer	20240824162544402.jpg
crystal	20240824162544402.jpg
hull	20240824162544402.jpg
woman	20240824162544402.jpg
shape	20240824162544402.jpg
imposing	20240824162544402.jpg
modern	20240824162544402.jpg
minimalism	20240824162544402.jpg
elegant	20240824162544402.jpg
dystopia	20240824162544402.jpg
mysterious	20240824162544402.jpg
godlike	20240824162544402.jpg
scary	20240824162544402.jpg
in	20240824162544402.jpg
dark	20240824162544402.jpg
room	20240824162544402.jpg
8k	20240824162544402.jpg
the	202408241625529530.jpg
pixelated	202408241625529530.jpg
beach	202408241625529530.jpg
scene	202408241625529530.jpg
with	202408241625529530.jpg
vibrant	202408241625529530.jpg
orange	202408241625529530.jpg
pink	202408241625529530.jpg
sunset	202408241625529530.jpg
silhouetted	202408241625529530.jpg
palm	202408241625529530.jpg
trees	202408241625529530.jpg
swaying	202408241625529530.jpg
in	202408241625529530.jpg
breeze	202408241625529530.jpg
is	202408241625529530.jpg
beautiful	202408241625529530.jpg
serene	202408241625529530.jpg
sight	202408241625529530.jpg
photo	202408241625591185.jpg
concept	202408241625591185.jpg
car	202408241625591185.jpg
from	202408241625591185.jpg
2025	202408241625591185.jpg
8k	202408241625591185.jpg
resoultion	202408241625591185.jpg
hyper	202408241625591185.jpg
realstic	202408241625591185.jpg
black	202408241625591185.jpg
night	202408241625591185.jpg
35mm	202408241625591185.jpg
film	202408241625591185.jpg
editorial	202408241625591185.jpg
high	202408241625591185.jpg
fashion	202408241625591185.jpg
need	202408241625591185.jpg
for	202408241625591185.jpg
speed	202408241625591185.jpg
cute	202408241626056496.jpg
grey	202408241626056496.jpg
cat	202408241626056496.jpg
in	202408241626056496.jpg
boat	202408241626056496.jpg
floating	202408241626056496.jpg
starry	202408241626056496.jpg
night	202408241626056496.jpg
sky	202408241626056496.jpg
over	202408241626056496.jpg
snowy	202408241626056496.jpg
city	202408241626056496.jpg
with	202408241626056496.jpg
little	202408241626056496.jpg
boy	202408241626056496.jpg
wearing	202408241626056496.jpg
pajamas	202408241626056496.jpg
studio	202408241626056496.jpg
ghibli	202408241626056496.jpg
style	202408241626056496.jpg
white	202408241626056496.jpg
background	202408241626056496.jpg
knight	20240824162614891.jpg
on	20240824162614891.jpg
snow	20240824162614891.jpg
covered	20240824162614891.jpg
mountain	20240824162614891.jpg
in	20240824162614891.jpg
front	20240824162614891.jpg
castle	20240824162614891.jpg
the	20240824162614891.jpg
style	20240824162614891.jpg
goblincore	20240824162614891.jpg
eroded	20240824162614891.jpg
interiors	20240824162614891.jpg
intel	20240824162614891.jpg
core	20240824162614891.jpg
dark	20240824162614891.jpg
cyan	20240824162614891.jpg
orange	20240824162614891.jpg
32k	20240824162614891.jpg
uhd	20240824162614891.jpg
made	20240824162614891.jpg
wrought	20240824162614891.jpg
iron	20240824162614891.jpg
heavy	20240824162614891.jpg
shading	20240824162614891.jpg
futuristic	20240824162630867.jpg
starship	20240824162630867.jpg
hyper	20240824162630867.jpg
realistic	20240824162630867.jpg
detailed	20240824162630867.jpg
render	20240824162630867.jpg
extremely	20240824162630867.jpg
complex	20240824162630867.jpg
advanced	20240824162630867.jpg
chassis	20240824162630867.jpg
natural	20240824162630867.jpg
dirt	20240824162630867.jpg
debris	20240824162630867.jpg
detail	20240824162630867.jpg
scuffs	20240824162630867.jpg
stunning	20240824162630867.jpg
details	20240824162630867.jpg
dark	20240824162630867.jpg
hues	20240824162630867.jpg
flying	20240824162630867.jpg
in	20240824162630867.jpg
the	20240824162630867.jpg
space	20240824162630867.jpg
high-tech	202408241626402517.jpg
government	202408241626402517.jpg
center	202408241626402517.jpg
country	202408241626402517.jpg
located	202408241626402517.jpg
in	202408241626402517.jpg
the	202408241626402517.jpg
on	202408241626402517.jpg
glacier	202408241626402517.jpg
with	202408241626402517.jpg
floating	202408241626402517.jpg
island	202408241626402517.jpg
used	202408241626402517.jpg
to	202408241626402517.jpg
store	202408241626402517.jpg
important	202408241626402517.jpg
documents	202408241626402517.jpg
holographic	202408241626402517.jpg
dome	202408241626402517.jpg
fire	202408241626402517.jpg
station	202408241626402517.jpg
police	202408241626402517.jpg
hospital	202408241626402517.jpg
train	202408241626402517.jpg
military	202408241626402517.jpg
base	202408241626402517.jpg
dragon	202408241626486334.jpg
on	202408241626486334.jpg
black	202408241626486334.jpg
background	202408241626486334.jpg
with	202408241626486334.jpg
sunrise	202408241626486334.jpg
in	202408241626486334.jpg
the	202408241626486334.jpg
silk	202408241626486334.jpg
screen	202408241626486334.jpg
by	202408241626486334.jpg
s	202408241626486334.jpg
ami	202408241626486334.jpg
reddit	202408241626486334.jpg
contest	202408241626486334.jpg
spring	202408241626486334.jpg
synchromism	202408241626486334.jpg
anime	202408241626486334.jpg
aesthetic	202408241626486334.jpg
sky	202408241626486334.jpg
vaporwave	202408241626486334.jpg
statue	202408241626581896.jpg
from	202408241626581896.jpg
muscular	202408241626581896.jpg
greek	202408241626581896.jpg
god	202408241626581896.jpg
with	202408241626581896.jpg
wavey	202408241626581896.jpg
middel	202408241626581896.jpg
parting	202408241626581896.jpg
hair	202408241626581896.jpg
out	202408241626581896.jpg
white	202408241626581896.jpg
marbel	202408241626581896.jpg
black	202408241626581896.jpg
backround	202408241626581896.jpg
standing	202408241626581896.jpg
on	202408241626581896.jpg
podest	202408241626581896.jpg
animeted	202408241626581896.jpg
style	202408241626581896.jpg
8k	202408241626581896.jpg
cube	202408241627497442.jpg
cutout	202408241627497442.jpg
an	202408241627497442.jpg
isometric	202408241627497442.jpg
programmer	202408241627497442.jpg
s	202408241627497442.jpg
bedroom	202408241627497442.jpg
3d	202408241627497442.jpg
art	202408241627497442.jpg
professional	202408241627497442.jpg
colors	202408241627497442.jpg
soft	202408241627497442.jpg
lighting	202408241627497442.jpg
high	202408241627497442.jpg
detail	202408241627497442.jpg
artstation	202408241627497442.jpg
concept	202408241627497442.jpg
behance	202408241627497442.jpg
ray	202408241627497442.jpg
tracing	202408241627497442.jpg
all	202408241627587662.jpg
glassblowing	202408241627587662.jpg
vampire	202408241627587662.jpg
urban	202408241627587662.jpg
style	202408241627587662.jpg
shooting	202408241627587662.jpg
stars	202408241627587662.jpg
big	202408241627587662.jpg
crescent	202408241627587662.jpg
moon	202408241627587662.jpg
one	202408241627587662.jpg
white	202408241627587662.jpg
wolf	202408241627587662.jpg
black	202408241627587662.jpg
howling	202408241627587662.jpg
at	202408241627587662.jpg
the	202408241627587662.jpg
constellations	202408241627587662.jpg
night	202408241627587662.jpg
spark	202408241627587662.jpg
blue	202408241627587662.jpg
light	202408241627587662.jpg
bats	202408241627587662.jpg
red	202408241627587662.jpg
purple	202408241627587662.jpg
cream	202408241627587662.jpg
wild	202408241628056788.jpg
west	202408241628056788.jpg
gunslinger	202408241628056788.jpg
male	202408241628056788.jpg
outlaw	202408241628056788.jpg
portrait	202408241628056788.jpg
studio	202408241628056788.jpg
ghibli	202408241628056788.jpg
akira	202408241628056788.jpg
toriyama	202408241628056788.jpg
james	202408241628056788.jpg
gilleard	202408241628056788.jpg
trending	202408241628056788.jpg
pixiv	202408241628056788.jpg
fanboxe	202408241628056788.jpg
4k	202408241628056788.jpg
the	202408241628056788.jpg
style	202408241628056788.jpg
90	202408241628056788.jpg
s	202408241628056788.jpg
vintage	202408241628056788.jpg
anime	202408241628056788.jpg
brown	202408241628056788.jpg
hair	202408241628056788.jpg
anime	20240824162826948.jpg
art	20240824162826948.jpg
akira	20240824162826948.jpg
detailed	20240824162826948.jpg
scene	20240824162826948.jpg
red	20240824162826948.jpg
perfect	20240824162826948.jpg
face	20240824162826948.jpg
intricately	20240824162826948.jpg
photorealism	20240824162826948.jpg
trending	20240824162826948.jpg
on	20240824162826948.jpg
artstation	20240824162826948.jpg
neon	20240824162826948.jpg
lights	20240824162826948.jpg
rainy	20240824162826948.jpg
day	20240824162826948.jpg
ray-traced	20240824162826948.jpg
environment	20240824162826948.jpg
vintage	20240824162826948.jpg
90	20240824162826948.jpg
s	20240824162826948.jpg
artwork	20240824162826948.jpg
the	202408241628342347.jpg
joker	202408241628342347.jpg
walking	202408241628342347.jpg
through	202408241628342347.jpg
streets	202408241628342347.jpg
new	202408241628342347.jpg
york	202408241628342347.jpg
stunning	202408241628342347.jpg
photo	202408241628342347.jpg
dark	202408241628342347.jpg
moody	202408241628342347.jpg
aesthetic	202408241628342347.jpg
at	202408241628342347.jpg
night	202408241628342347.jpg
city	202408241628342347.jpg
lights	202408241628342347.jpg
in	202408241628342347.jpg
background	202408241628342347.jpg
surreal	202408241628342347.jpg
8k	202408241628342347.jpg
shabby	202408241628447113.jpg
chic	202408241628447113.jpg
dreamy	202408241628447113.jpg
mist	202408241628447113.jpg
pastel	202408241628447113.jpg
junk	202408241628447113.jpg
journals	202408241628447113.jpg
christmas	202408241628447113.jpg
street	202408241628447113.jpg
with	202408241628447113.jpg
cafe	202408241628447113.jpg
shops	202408241628447113.jpg
swirling	202408241628447113.jpg
magical	202408241628447113.jpg
fairytale	202408241628447113.jpg
abstract	202408241628447113.jpg
art	202408241628447113.jpg
style	202408241628447113.jpg
warhammer	202408241628518681.jpg
40k	202408241628518681.jpg
chaos	202408241628518681.jpg
armor	202408241628518681.jpg
black	202408241628518681.jpg
hair	202408241628518681.jpg
jaw	202408241628518681.jpg
made	202408241628518681.jpg
from	202408241628518681.jpg
metal	202408241628518681.jpg
holding	202408241628518681.jpg
bolt	202408241628518681.jpg
gun	202408241628518681.jpg
vertical	202408241628584219.jpg
internal	202408241628584219.jpg
circular	202408241628584219.jpg
cylindre	202408241628584219.jpg
corridor	202408241628584219.jpg
with	202408241628584219.jpg
minimalist	202408241628584219.jpg
space	202408241628584219.jpg
an	202408241628584219.jpg
ogive	202408241628584219.jpg
roof	202408241628584219.jpg
white	202408241628584219.jpg
sand	202408241628584219.jpg
beautiful	202408241629068500.jpg
illustration	202408241629068500.jpg
man	202408241629068500.jpg
woman	202408241629068500.jpg
in	202408241629068500.jpg
pink	202408241629068500.jpg
dress	202408241629068500.jpg
watching	202408241629068500.jpg
the	202408241629068500.jpg
sunset	202408241629068500.jpg
setting	202408241629068500.jpg
over	202408241629068500.jpg
golden	202408241629068500.jpg
gate	202408241629068500.jpg
bridge	202408241629068500.jpg
hour	202408241629068500.jpg
stunning	202408241629068500.jpg
color	202408241629068500.jpg
scheme	202408241629068500.jpg
masterpiece	202408241629068500.jpg
stunning	202408241629149280.jpg
iridescent	202408241629149280.jpg
male	202408241629149280.jpg
marble	202408241629149280.jpg
bust	202408241629149280.jpg
very	202408241629149280.jpg
reflective	202408241629149280.jpg
perfect	202408241629149280.jpg
lighting	202408241629149280.jpg
dark	202408241629149280.jpg
background	202408241629149280.jpg
zbrush	202408241629149280.jpg
realistic	202408241629217321.jpg
painting	202408241629217321.jpg
stack	202408241629217321.jpg
pancakes	202408241629217321.jpg
bacon	202408241629217321.jpg
eggs	202408241629217321.jpg
with	202408241629217321.jpg
complementary	202408241629217321.jpg
colors	202408241629217321.jpg
beautiful	202408241629272060.jpg
cozy	202408241629272060.jpg
bedroom	202408241629272060.jpg
with	202408241629272060.jpg
floor	202408241629272060.jpg
to	202408241629272060.jpg
ceiling	202408241629272060.jpg
glass	202408241629272060.jpg
windows	202408241629272060.jpg
overlooking	202408241629272060.jpg
cyberpunk	202408241629272060.jpg
city	202408241629272060.jpg
at	202408241629272060.jpg
night	202408241629272060.jpg
thunderstorm	202408241629272060.jpg
outside	202408241629272060.jpg
torrential	202408241629272060.jpg
rain	202408241629272060.jpg
detailed	202408241629272060.jpg
high	202408241629272060.jpg
resolution	202408241629272060.jpg
photorrealistic	202408241629272060.jpg
dark	202408241629272060.jpg
gloomy	202408241629272060.jpg
moody	202408241629272060.jpg
aesthetic	202408241629272060.jpg
create	202408241629372780.jpg
oil	202408241629372780.jpg
painting	202408241629372780.jpg
surrealistic	202408241629372780.jpg
in	202408241629372780.jpg
michael	202408241629372780.jpg
parkes	202408241629372780.jpg
style	202408241629372780.jpg
raspberries	202408241629372780.jpg
cherries	202408241629372780.jpg
on	202408241629372780.jpg
cheese	202408241629372780.jpg
cake	202408241629372780.jpg
with	202408241629372780.jpg
marmelade	202408241629372780.jpg
background	202408241629372780.jpg
pale	202408241629372780.jpg
blue	202408241629372780.jpg
gold	202408241629372780.jpg
geometric	202408241629372780.jpg
pattern	202408241629372780.jpg
like	202408241629372780.jpg
vasily	202408241629372780.jpg
kandinsky	202408241629372780.jpg
composition	202408241629372780.jpg
8	202408241629372780.jpg
please	202408241629372780.jpg
no	202408241629372780.jpg
woman	202408241629372780.jpg
faces	202408241629372780.jpg
only	202408241629372780.jpg
berries	202408241629372780.jpg
stilllife	202408241629372780.jpg
gustav	202408241629372780.jpg
klimt	202408241629372780.jpg
john	202408241629372780.jpg
william	202408241629372780.jpg
waterhouse	202408241629372780.jpg
very	202408241629372780.jpg
high	202408241629372780.jpg
resolution	202408241629372780.jpg
realistic	202408241629372780.jpg
40	202408241629372780.jpg
k	202408241629372780.jpg
detailed	202408241629432651.jpg
vibrant	202408241629432651.jpg
illustration	202408241629432651.jpg
cowboy	202408241629432651.jpg
in	202408241629432651.jpg
the	202408241629432651.jpg
copper	202408241629432651.jpg
canyons	202408241629432651.jpg
sierra	202408241629432651.jpg
chihuahua	202408241629432651.jpg
state	202408241629432651.jpg
by	202408241629432651.jpg
herge	202408241629432651.jpg
style	202408241629432651.jpg
tin-tin	202408241629432651.jpg
comics	202408241629432651.jpg
colors	202408241629432651.jpg
sunny	202408241629432651.jpg
day	202408241629432651.jpg
attention	202408241629432651.jpg
to	202408241629432651.jpg
detail	202408241629432651.jpg
8k	202408241629432651.jpg
sphere	202408241629506434.jpg
shaped	202408241629506434.jpg
like	202408241629506434.jpg
drought	202408241629506434.jpg
with	202408241629506434.jpg
the	202408241629506434.jpg
lower	202408241629506434.jpg
part	202408241629506434.jpg
resembling	202408241629506434.jpg
water	202408241629506434.jpg
dripping	202408241629506434.jpg
drop	202408241629506434.jpg
falls	202408241629506434.jpg
from	202408241629506434.jpg
set	202408241629506434.jpg
against	202408241629506434.jpg
blue	202408241629506434.jpg
sky	202408241629506434.jpg
white	202408241629506434.jpg
clouds	202408241629506434.jpg
surrealism	202408241629506434.jpg
art	202408241629506434.jpg
back	202408241630061375.jpg
vew	202408241630061375.jpg
mage	202408241630061375.jpg
man	202408241630061375.jpg
walking	202408241630061375.jpg
in	202408241630061375.jpg
cobbeled	202408241630061375.jpg
street	202408241630061375.jpg
medieval	202408241630061375.jpg
village	202408241630061375.jpg
covered	202408241630061375.jpg
snow	202408241630061375.jpg
during	202408241630061375.jpg
the	202408241630061375.jpg
night	202408241630061375.jpg
he	202408241630061375.jpg
wearing	202408241630061375.jpg
hooded	202408241630061375.jpg
coat	202408241630061375.jpg
futuristic	2024082416301216.jpg
sci-fi	2024082416301216.jpg
pod	2024082416301216.jpg
chair	2024082416301216.jpg
flat	2024082416301216.jpg
design	2024082416301216.jpg
product-view	2024082416301216.jpg
editorial	2024082416301216.jpg
photography	2024082416301216.jpg
transparent	2024082416301216.jpg
orb	2024082416301216.jpg
product	2024082416301216.jpg
natural	2024082416301216.jpg
lighting	2024082416301216.jpg
plants	2024082416301216.jpg
daytime	2024082416301216.jpg
zbrush	2024082416301216.jpg
8k	2024082416301216.jpg
wooden	2024082416301216.jpg
environment	2024082416301216.jpg
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, password_hash) FROM stdin;
1	Anonym	617367645b6a616f7071713d756a6831325d34363d662ec7511d1fabd23636d2294d366ced250f959adc
2	Admin	617367645b6a616f7071713d756a6831325d34363d662ec7511d1fabd23636d2294d366ced250f959adc
3	test	617367645b6a616f7071713d756a6831325d34363d667c4a8d09ca3762af61e59520943dc26494f8941b
4	test1	617367645b6a616f7071713d756a6831325d34363d6640bd001563085fc35165329ea1ff5c5ecbdbbeef
\.


--
-- Name: consideration_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.consideration_id_seq', 40, true);


--
-- Name: images_image_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.images_image_id_seq', 40, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 4, true);


--
-- Name: consideration consideration_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.consideration
    ADD CONSTRAINT consideration_pkey PRIMARY KEY (id);


--
-- Name: images images_image_url_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.images
    ADD CONSTRAINT images_image_url_key UNIQUE (image_url);


--
-- Name: images images_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.images
    ADD CONSTRAINT images_pkey PRIMARY KEY (user_id, image_url);


--
-- Name: likes likes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.likes
    ADD CONSTRAINT likes_pkey PRIMARY KEY (user_id, image_url);


--
-- Name: promts promts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.promts
    ADD CONSTRAINT promts_pkey PRIMARY KEY (title, image_url);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: consideration consideration_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.consideration
    ADD CONSTRAINT consideration_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: images images_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.images
    ADD CONSTRAINT images_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: likes likes_image_url_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.likes
    ADD CONSTRAINT likes_image_url_fkey FOREIGN KEY (image_url) REFERENCES public.images(image_url) ON DELETE CASCADE;


--
-- Name: likes likes_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.likes
    ADD CONSTRAINT likes_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: promts promts_image_url_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.promts
    ADD CONSTRAINT promts_image_url_fkey FOREIGN KEY (image_url) REFERENCES public.images(image_url) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

