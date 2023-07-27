--
-- PostgreSQL database dump
--

-- Dumped from database version 14.3
-- Dumped by pg_dump version 14.3

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
-- Name: Address; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Address" (
    id integer NOT NULL,
    street character varying NOT NULL,
    house character varying NOT NULL
);


ALTER TABLE public."Address" OWNER TO postgres;

--
-- Name: Address_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Address_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Address_id_seq" OWNER TO postgres;

--
-- Name: Address_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Address_id_seq" OWNED BY public."Address".id;


--
-- Name: Deposit; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Deposit" (
    id integer NOT NULL,
    user_id bigint NOT NULL,
    amount double precision NOT NULL,
    date character varying NOT NULL
);


ALTER TABLE public."Deposit" OWNER TO postgres;

--
-- Name: Deposits_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Deposits_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Deposits_id_seq" OWNER TO postgres;

--
-- Name: Deposits_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Deposits_id_seq" OWNED BY public."Deposit".id;


--
-- Name: Expense; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Expense" (
    id integer NOT NULL,
    user_id integer NOT NULL,
    amount double precision NOT NULL,
    service character varying NOT NULL,
    date character varying NOT NULL
);


ALTER TABLE public."Expense" OWNER TO postgres;

--
-- Name: Expenses_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Expenses_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Expenses_id_seq" OWNER TO postgres;

--
-- Name: Expenses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Expenses_id_seq" OWNED BY public."Expense".id;


--
-- Name: Faq; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Faq" (
    id integer NOT NULL,
    question character varying NOT NULL,
    answer character varying NOT NULL
);


ALTER TABLE public."Faq" OWNER TO postgres;

--
-- Name: Faq_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Faq_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Faq_id_seq" OWNER TO postgres;

--
-- Name: Faq_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Faq_id_seq" OWNED BY public."Faq".id;


--
-- Name: Role; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Role" (
    id integer NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public."Role" OWNER TO postgres;

--
-- Name: Role_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Role_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Role_id_seq" OWNER TO postgres;

--
-- Name: Role_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Role_id_seq" OWNED BY public."Role".id;


--
-- Name: SEO; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."SEO" (
    id integer NOT NULL,
    title character varying NOT NULL,
    keywords character varying NOT NULL,
    description character varying NOT NULL,
    uri character varying NOT NULL
);


ALTER TABLE public."SEO" OWNER TO postgres;

--
-- Name: SEO_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."SEO_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."SEO_id_seq" OWNER TO postgres;

--
-- Name: SEO_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."SEO_id_seq" OWNED BY public."SEO".id;


--
-- Name: Service; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Service" (
    id integer NOT NULL,
    name character varying NOT NULL,
    note character varying,
    full_price double precision,
    rent_price double precision,
    type_id integer NOT NULL
);


ALTER TABLE public."Service" OWNER TO postgres;

--
-- Name: Service_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Service_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Service_id_seq" OWNER TO postgres;

--
-- Name: Service_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Service_id_seq" OWNED BY public."Service".id;


--
-- Name: Service_type; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Service_type" (
    id integer NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public."Service_type" OWNER TO postgres;

--
-- Name: Service_type_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Service_type_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Service_type_id_seq" OWNER TO postgres;

--
-- Name: Service_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Service_type_id_seq" OWNED BY public."Service_type".id;


--
-- Name: Session; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Session" (
    hash character varying NOT NULL,
    user_id bigint NOT NULL,
    date time without time zone NOT NULL
);


ALTER TABLE public."Session" OWNER TO postgres;

--
-- Name: Settings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Settings" (
    id integer NOT NULL,
    key character varying NOT NULL,
    value character varying NOT NULL,
    description character varying NOT NULL
);


ALTER TABLE public."Settings" OWNER TO postgres;

--
-- Name: Settings_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Settings_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Settings_id_seq" OWNER TO postgres;

--
-- Name: Settings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Settings_id_seq" OWNED BY public."Settings".id;


--
-- Name: Tariff; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Tariff" (
    id integer NOT NULL,
    type_id integer NOT NULL,
    price double precision NOT NULL,
    name character varying NOT NULL,
    description character varying NOT NULL,
    speed integer,
    digital_channel integer,
    analog_channel integer,
    image character varying,
    color character varying NOT NULL
);


ALTER TABLE public."Tariff" OWNER TO postgres;

--
-- Name: Tariff_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Tariff_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Tariff_id_seq" OWNER TO postgres;

--
-- Name: Tariff_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Tariff_id_seq" OWNED BY public."Tariff".id;


--
-- Name: Tariff_type; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Tariff_type" (
    id integer NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public."Tariff_type" OWNER TO postgres;

--
-- Name: Type_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Type_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Type_id_seq" OWNER TO postgres;

--
-- Name: Type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Type_id_seq" OWNED BY public."Tariff_type".id;


--
-- Name: User; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."User" (
    id integer NOT NULL,
    name character varying NOT NULL,
    phone character varying NOT NULL,
    account_number character varying NOT NULL,
    password character varying NOT NULL,
    current_balance double precision NOT NULL,
    current_tariff bigint,
    role integer NOT NULL,
    house integer NOT NULL,
    flat integer,
    baned integer NOT NULL
);


ALTER TABLE public."User" OWNER TO postgres;

--
-- Name: User_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."User_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."User_id_seq" OWNER TO postgres;

--
-- Name: User_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."User_id_seq" OWNED BY public."User".id;


--
-- Name: Address id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Address" ALTER COLUMN id SET DEFAULT nextval('public."Address_id_seq"'::regclass);


--
-- Name: Deposit id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Deposit" ALTER COLUMN id SET DEFAULT nextval('public."Deposits_id_seq"'::regclass);


--
-- Name: Expense id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Expense" ALTER COLUMN id SET DEFAULT nextval('public."Expenses_id_seq"'::regclass);


--
-- Name: Faq id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Faq" ALTER COLUMN id SET DEFAULT nextval('public."Faq_id_seq"'::regclass);


--
-- Name: Role id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Role" ALTER COLUMN id SET DEFAULT nextval('public."Role_id_seq"'::regclass);


--
-- Name: SEO id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."SEO" ALTER COLUMN id SET DEFAULT nextval('public."SEO_id_seq"'::regclass);


--
-- Name: Service id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Service" ALTER COLUMN id SET DEFAULT nextval('public."Service_id_seq"'::regclass);


--
-- Name: Service_type id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Service_type" ALTER COLUMN id SET DEFAULT nextval('public."Service_type_id_seq"'::regclass);


--
-- Name: Settings id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Settings" ALTER COLUMN id SET DEFAULT nextval('public."Settings_id_seq"'::regclass);


--
-- Name: Tariff id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Tariff" ALTER COLUMN id SET DEFAULT nextval('public."Tariff_id_seq"'::regclass);


--
-- Name: Tariff_type id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Tariff_type" ALTER COLUMN id SET DEFAULT nextval('public."Type_id_seq"'::regclass);


--
-- Name: User id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."User" ALTER COLUMN id SET DEFAULT nextval('public."User_id_seq"'::regclass);


--
-- Data for Name: Address; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Address" (id, street, house) FROM stdin;
1	Пушкина	72
2	Красина	66
3	Гоголя	12
4	Кирова	23
6	Пушкина	22
7	3 микрорайон	16
8	Пичугина	16
\.


--
-- Data for Name: Deposit; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Deposit" (id, user_id, amount, date) FROM stdin;
4	6	200	2023-06-06 15:36:42.7917836+05:00
5	6	100	2023-06-06 15:48:23.7518713+05:00
6	67	700	2023-06-16 00:31:03.1794828+05:00
7	6	100	2023-06-16 02:20:54.8127538+05:00
8	6	500	2023-06-16 02:45:34.736733+05:00
9	6	100	2023-06-16 03:18:23.4883158+05:00
10	6	100	2023-06-16 08:36:52.1249777+05:00
11	6	200	2023-06-16 09:39:08.5162149+05:00
\.


--
-- Data for Name: Expense; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Expense" (id, user_id, amount, service, date) FROM stdin;
\.


--
-- Data for Name: Faq; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Faq" (id, question, answer) FROM stdin;
2	Какой тип подключения используется	В сети Орбител используется статический IP.
3	Когда списывается абонентская плата	Абонентская плата списывается в начале каждого месяца независимо от даты вашего подключения. Списание производится за весь месяц целиком. Если вы были подключены в середине месяца, первый раз вам начисляется абонентская плата на остаток месяца, затем будет начисляться за полный.
4	Что делать если пропал ТВ-канал	Если пропал ТВ-канал попробуйте выполнить автонастройку каналов <a>по инструкции.<a>http://192.168.0.105:8080/tv-manual
5	Что делать если я переезжаю?	В этом случае вам необходимо обратиться в офис или по телефону и уточнить возможность переоформления договора по новому адресу.
6	Как заблокировать услугу?	Если вы уезжаете в отпуск и не планируете пользоваться услугой, вы можете подойти в офис или подать заявку по телефону на блокировку.
7	Почему у меня на счету минус?	По условию договора абонент обязуется оплачивать абонентскую плату ежемесячно. Если вы не планируете пользоваться интернетом какое то время, вам нужно подать заявку на блокировку. Но если задолженность все-таки образовалась, а интернетом в течении месяца вы не пользовались, вам могут сделать перерасчет за данный период, удержав с вас по 100 рублей за поддержание линии.
8	Я прописал настройки на другом компьютере, а интернета нет	Если вы прописали настройки на другом компьютере или установили роутер, вам необходимо позвонить в тех. поддержку по тел. 8 (3522) 65-00-00, чтобы ваше устройство "зафиксировали" в сети Орбител.
1	Техническая поддержка	В компании Орбител работает круглосуточная тех. поддержка, тел. 8 (3522) 65-00-00.
11	фыв	123
\.


--
-- Data for Name: Role; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Role" (id, name) FROM stdin;
1	Admin
2	Moderator
3	Default
\.


--
-- Data for Name: SEO; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."SEO" (id, title, keywords, description, uri) FROM stdin;
1	Orbitel - интернет, кабельное ТВ, телефония	интернет, кабельное ТВ, телефония	ООО «Орбител» — оператор связи, оказывающий широкий спектр качественных услуг физическим и юридическим лицам г.Кургана.  Компания имеет собственную волоконно-оптическую сеть, в опорных узлах которой установлено высокопроизводительное и надежное оборудование фирмы Cisco. «Городская сеть передачи данных» ООО «Орбител» построена по технологии «оптика до дома» (FTTB), что обеспечивает отличные показатели и высокую надежность оказываемых услуг.	/
4	Телевидение - Orbitel	интернет, кабельное ТВ, телефония	ООО «Орбител» — оператор связи, оказывающий широкий спектр качественных услуг физическим и юридическим лицам г.Кургана.  Компания имеет собственную волоконно-оптическую сеть, в опорных узлах которой установлено высокопроизводительное и надежное оборудование фирмы Cisco. «Городская сеть передачи данных» ООО «Орбител» построена по технологии «оптика до дома» (FTTB), что обеспечивает отличные показатели и высокую надежность оказываемых услуг.	/tv
8	Телефон для ЮЛ - Orbitel	интернет, кабельное ТВ, телефония	ООО «Орбител» — оператор связи, оказывающий широкий спектр качественных услуг физическим и юридическим лицам г.Кургана.  Компания имеет собственную волоконно-оптическую сеть, в опорных узлах которой установлено высокопроизводительное и надежное оборудование фирмы Cisco. «Городская сеть передачи данных» ООО «Орбител» построена по технологии «оптика до дома» (FTTB), что обеспечивает отличные показатели и высокую надежность оказываемых услуг.	/phone-for-ul
19	Личный кабинет - Orbitel	интернет, кабельное ТВ, телефония	ООО «Орбител» — оператор связи, оказывающий широкий спектр качественных услуг физическим и юридическим лицам г.Кургана.  Компания имеет собственную волоконно-оптическую сеть, в опорных узлах которой установлено высокопроизводительное и надежное оборудование фирмы Cisco. «Городская сеть передачи данных» ООО «Орбител» построена по технологии «оптика до дома» (FTTB), что обеспечивает отличные показатели и высокую надежность оказываемых услуг.	/personal_account
18	О нас - Orbitel	интернет, кабельное ТВ, телефония	ООО «Орбител» — оператор связи, оказывающий широкий спектр качественных услуг физическим и юридическим лицам г.Кургана.  Компания имеет собственную волоконно-оптическую сеть, в опорных узлах которой установлено высокопроизводительное и надежное оборудование фирмы Cisco. «Городская сеть передачи данных» ООО «Орбител» построена по технологии «оптика до дома» (FTTB), что обеспечивает отличные показатели и высокую надежность оказываемых услуг.	/about
2	Интернет - Orbitel	интернет, кабельное ТВ, телефония	ООО «Орбител» — оператор связи, оказывающий широкий спектр качественных услуг физическим и юридическим лицам г.Кургана.  Компания имеет собственную волоконно-оптическую сеть, в опорных узлах которой установлено высокопроизводительное и надежное оборудование фирмы Cisco. «Городская сеть передачи данных» ООО «Орбител» построена по технологии «оптика до дома» (FTTB), что обеспечивает отличные показатели и высокую надежность оказываемых услуг.	/tarif_for_home
5	Телефон - Orbitel	интернет, кабельное ТВ, телефония	ООО «Орбител» — оператор связи, оказывающий широкий спектр качественных услуг физическим и юридическим лицам г.Кургана.  Компания имеет собственную волоконно-оптическую сеть, в опорных узлах которой установлено высокопроизводительное и надежное оборудование фирмы Cisco. «Городская сеть передачи данных» ООО «Орбител» построена по технологии «оптика до дома» (FTTB), что обеспечивает отличные показатели и высокую надежность оказываемых услуг.	/phone
14	Способы оплаты - Orbitel	интернет, кабельное ТВ, телефония	ООО «Орбител» — оператор связи, оказывающий широкий спектр качественных услуг физическим и юридическим лицам г.Кургана.  Компания имеет собственную волоконно-оптическую сеть, в опорных узлах которой установлено высокопроизводительное и надежное оборудование фирмы Cisco. «Городская сеть передачи данных» ООО «Орбител» построена по технологии «оптика до дома» (FTTB), что обеспечивает отличные показатели и высокую надежность оказываемых услуг.	/oplata
6	Дополнительные услуги для ДОМА - Orbitel	интернет, кабельное ТВ, телефония	ООО «Орбител» — оператор связи, оказывающий широкий спектр качественных услуг физическим и юридическим лицам г.Кургана.  Компания имеет собственную волоконно-оптическую сеть, в опорных узлах которой установлено высокопроизводительное и надежное оборудование фирмы Cisco. «Городская сеть передачи данных» ООО «Орбител» построена по технологии «оптика до дома» (FTTB), что обеспечивает отличные показатели и высокую надежность оказываемых услуг.	/dop_for_home
12	Настройка роутера - Orbitel	интернет, кабельное ТВ, телефония	ООО «Орбител» — оператор связи, оказывающий широкий спектр качественных услуг физическим и юридическим лицам г.Кургана.  Компания имеет собственную волоконно-оптическую сеть, в опорных узлах которой установлено высокопроизводительное и надежное оборудование фирмы Cisco. «Городская сеть передачи данных» ООО «Орбител» построена по технологии «оптика до дома» (FTTB), что обеспечивает отличные показатели и высокую надежность оказываемых услуг.	/routers
11	Дополнительные услуги для БИЗНЕСА - Orbitel	интернет, кабельное ТВ, телефония	ООО «Орбител» — оператор связи, оказывающий широкий спектр качественных услуг физическим и юридическим лицам г.Кургана.  Компания имеет собственную волоконно-оптическую сеть, в опорных узлах которой установлено высокопроизводительное и надежное оборудование фирмы Cisco. «Городская сеть передачи данных» ООО «Орбител» построена по технологии «оптика до дома» (FTTB), что обеспечивает отличные показатели и высокую надежность оказываемых услуг.	/dop-for-ul
10	Виртуальная АТС - Orbitel	интернет, кабельное ТВ, телефония	ООО «Орбител» — оператор связи, оказывающий широкий спектр качественных услуг физическим и юридическим лицам г.Кургана.  Компания имеет собственную волоконно-оптическую сеть, в опорных узлах которой установлено высокопроизводительное и надежное оборудование фирмы Cisco. «Городская сеть передачи данных» ООО «Орбител» построена по технологии «оптика до дома» (FTTB), что обеспечивает отличные показатели и высокую надежность оказываемых услуг.	/virtual-ats
9	Каналы передачи данных для ЮЛ - Orbitel	интернет, кабельное ТВ, телефония	ООО «Орбител» — оператор связи, оказывающий широкий спектр качественных услуг физическим и юридическим лицам г.Кургана.  Компания имеет собственную волоконно-оптическую сеть, в опорных узлах которой установлено высокопроизводительное и надежное оборудование фирмы Cisco. «Городская сеть передачи данных» ООО «Орбител» построена по технологии «оптика до дома» (FTTB), что обеспечивает отличные показатели и высокую надежность оказываемых услуг.	/vpn-for-ul
3	Интернет + ТВ - Orbitel	интернет, кабельное ТВ, телефония	ООО «Орбител» — оператор связи, оказывающий широкий спектр качественных услуг физическим и юридическим лицам г.Кургана.  Компания имеет собственную волоконно-оптическую сеть, в опорных узлах которой установлено высокопроизводительное и надежное оборудование фирмы Cisco. «Городская сеть передачи данных» ООО «Орбител» построена по технологии «оптика до дома» (FTTB), что обеспечивает отличные показатели и высокую надежность оказываемых услуг.	/inettv
13	Настройка ТВ каналов - Orbitel	интернет, кабельное ТВ, телефония	ООО «Орбител» — оператор связи, оказывающий широкий спектр качественных услуг физическим и юридическим лицам г.Кургана.  Компания имеет собственную волоконно-оптическую сеть, в опорных узлах которой установлено высокопроизводительное и надежное оборудование фирмы Cisco. «Городская сеть передачи данных» ООО «Орбител» построена по технологии «оптика до дома» (FTTB), что обеспечивает отличные показатели и высокую надежность оказываемых услуг.	/tv-manual
15	Вопрос - ответ - Orbitel	интернет, кабельное ТВ, телефония	ООО «Орбител» — оператор связи, оказывающий широкий спектр качественных услуг физическим и юридическим лицам г.Кургана.  Компания имеет собственную волоконно-оптическую сеть, в опорных узлах которой установлено высокопроизводительное и надежное оборудование фирмы Cisco. «Городская сеть передачи данных» ООО «Орбител» построена по технологии «оптика до дома» (FTTB), что обеспечивает отличные показатели и высокую надежность оказываемых услуг.	/faq
16	Калькулятор тарифов - Orbitel	интернет, кабельное ТВ, телефония	ООО «Орбител» — оператор связи, оказывающий широкий спектр качественных услуг физическим и юридическим лицам г.Кургана.  Компания имеет собственную волоконно-оптическую сеть, в опорных узлах которой установлено высокопроизводительное и надежное оборудование фирмы Cisco. «Городская сеть передачи данных» ООО «Орбител» построена по технологии «оптика до дома» (FTTB), что обеспечивает отличные показатели и высокую надежность оказываемых услуг.	/calculator
17	Оформление заявки - Orbitel	интернет, кабельное ТВ, телефония	ООО «Орбител» — оператор связи, оказывающий широкий спектр качественных услуг физическим и юридическим лицам г.Кургана.  Компания имеет собственную волоконно-оптическую сеть, в опорных узлах которой установлено высокопроизводительное и надежное оборудование фирмы Cisco. «Городская сеть передачи данных» ООО «Орбител» построена по технологии «оптика до дома» (FTTB), что обеспечивает отличные показатели и высокую надежность оказываемых услуг.	/abonent_application
20	Авторизация - Orbitel	интернет, кабельное ТВ, телефония	ООО «Орбител» — оператор связи, оказывающий широкий спектр качественных услуг физическим и юридическим лицам г.Кургана.  Компания имеет собственную волоконно-оптическую сеть, в опорных узлах которой установлено высокопроизводительное и надежное оборудование фирмы Cisco. «Городская сеть передачи данных» ООО «Орбител» построена по технологии «оптика до дома» (FTTB), что обеспечивает отличные показатели и высокую надежность оказываемых услуг.	/authorization
7	Интернет для ЮЛ - Orbitel	интернет, кабельное ТВ, телефония	ООО «Орбител» — оператор связи, оказывающий широкий спектр качественных услуг физическим и юридическим лицам г.Кургана.  Компания имеет собственную волоконно-оптическую сеть, в опорных узлах которой установлено высокопроизводительное и надежное оборудование фирмы Cisco. «Городская сеть передачи данных» ООО «Орбител» построена по технологии «оптика до дома» (FTTB), что обеспечивает отличные показатели и высокую надежность оказываемых услуг.	/inet-for-ul
21	Пополнение баланса - Orbitel	интернет, кабельное ТВ, телефония	ООО «Орбител» — оператор связи, оказывающий широкий спектр качественных услуг физическим и юридическим лицам г.Кургана.  Компания имеет собственную волоконно-оптическую сеть, в опорных узлах которой установлено высокопроизводительное и надежное оборудование фирмы Cisco. «Городская сеть передачи данных» ООО «Орбител» построена по технологии «оптика до дома» (FTTB), что обеспечивает отличные показатели и высокую надежность оказываемых услуг.	/replenishment_balance
\.


--
-- Data for Name: Service; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Service" (id, name, note, full_price, rent_price, type_id) FROM stdin;
2	Ресивер World Vision T62D (DVB-C, DVB-T/T2)		1500	100	1
3	IP TV-приставка NV-501Wac	\N	4300	100	1
5	Настройка Wi-Fi роутера абонента в офисе, приобретенного у сторонней фирмы	\N	350	\N	2
6	Настройка Wi-Fi роутера абонента в офисе, приобретенного в компании «Орбител»	\N	0	\N	2
7	Настройка Wi-Fi роутера абонента на дому, приобретенного у сторонней фирмы	\N	600	\N	2
8	Настройка Wi-Fi роутера абонента на дому, приобретенного в компании «Орбител»	\N	250	\N	2
9	Консультация специалиста с выездом к Заказчику, за 1 час	\N	250	\N	2
10	Смена тарифного плана	\N	0	\N	3
11	Перезаключение договора по инициативе Оператора	\N	0	\N	3
12	Перезаключение договора по инициативе Абонента	\N	0	\N	3
13	Блокировка (ограничение доступа) по заявлению Абонента возможна 2 раза в год, суммарно не превышающих 6 мес.	в течение 3 рабочих дней с момента получения заявки Абонента	0	\N	3
14	Блокировка (ограничение доступа) по заявлению Абонента сроком более 6 мес.	в течение 3 рабочих дней с момента получения заявки Абонента	100	\N	3
15	Разблокировка (возобновление доступа) по заявке Абонента, поданной в срок не более 3-х месяцев с момента блокировки	в течение 3 рабочих дней с момента получения заявки Абонента	0	\N	3
16	Разблокировка (возобновление доступа) по заявке Абонента, поданной в срок более 3-х месяцев с момента блокировки	за каждый полный месяц блокировки сверх 3-х месяцев	0	\N	3
17	Детализация статистики по заявке Абонента за предыдущий месяц	в течение 5 рабочих дней с момента получения заявки Абонента	0	\N	3
18	Детализация статистики по заявке Абонента ранее предыдущего месяца	за 1 стандартную страницу формата А4, производится в течение 5 рабочих дней с момента получения заявки Абонента	10	\N	3
1	Wi-Fi-роутер (с настройкой)	    	1800	100	1
\.


--
-- Data for Name: Service_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Service_type" (id, name) FROM stdin;
1	Оборудование
2	Настройка роутера
3	Дополнительные услуги тарифного плана
4	Абонентская плата
\.


--
-- Data for Name: Session; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Session" (hash, user_id, date) FROM stdin;
f1d47e5f5497b613ab6d6110726dd76a8910b56a858230a1d3fca3e27040c279	1	10:25:58.714513
157a3ac39bc20e4ec3b50a0c5510106646457f5306131ed3c750693436b39aca	1	21:06:34.531546
fe268bf10f435cce34a7ff7837af7ade488e082f2b071380a533231a8e7bc6e4	1	21:12:00.166678
8522428781a8af59e78ba593a051f51e4f437b67c0a597b4cd71257963528289	1	23:30:15.115084
7dc0bb9a1bafb737b71284e0c6aa90158dc8d13e5f01631a7b49219d01e3cc05	1	12:09:35.576532
167a0028a8f7419c3869727da65ae1561f172240cac38e2117e4902f8aa0732f	1	02:46:42.258265
10f20eb1b99f69d7d24ad0d6a985d1b02980c5226ef2643e50b1febf71121b98	1	03:29:39.613342
d178c6d586ef2acff2b89b0243394f271cecc7cf4856787b42ffb2c91271e314	1	16:56:43.837663
d48984308808d25012b7ea9931d21e24868665f964bd985c9a26437c62adb529	1	14:42:43.59981
06d131fb4aa7843a30336e6451aba58c1761961bca2c50835e9a69a5a0fe23ac	1	16:09:27.143118
4cf6d0f8e6483a903c979516fc7c4ef55b335001d1c604cd35f11818ee33e200	1	00:33:32.242644
f682589ddea62d05cac1654175bad355f7fb967b0c0bae1d4569d627a423c452	1	07:01:04.156121
2c56d22db930df2006f8abcd5e5e7f95c7802c97a327a6301f03c7d606974e41	1	08:34:44.190328
f0eb2057358d8d760efaaf055f45422c1da76417d2b97343cb03f28f5e692ad3	6	15:30:25.319542
6914b4c69f739087c50b9b922e588ca8a9477c9a8a94f36080f003b4c23164b1	1	18:05:25.334711
5de6741c3d2680b8a57cdd5869ce06667626d2678ed60bdf6a7a7bfd6d9e8db6	1	18:05:33.072599
669127b57527fb572166f49d31380068d76523abafc73d903461dfac57a29f5a	1	23:05:46.630108
e0dd64a4e6acb8b2e66f0eb7d453ae51f7f1eea7678c44853a910b11e9bb19ef	1	23:08:11.118534
\.


--
-- Data for Name: Settings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Settings" (id, key, value, description) FROM stdin;
5	positive_massage_check-connection	Поздравляем! <span>Данный<span> адрес доступен для подключения!	 1
20	section_advantages_inet-for-ul_elements	Высокую надежность соединения 24/7;Скорость доступа до 1 Гб/сек;Грамотную и своевременную техническую поддержку;Кредитную систему оплаты	 1
1	sections_tariffs_title	Выберите тариф для подключения!!!	Заголовок секций с выбором тарифов
2	section_check-connection_title	Узнайте наличие подключения по вашему адресу!	Заголовк секции с проверкой подключения
7	negative_massage_check-connection	К сожалению, в данный момент<span> адрес не доступен для подключения.Приносим свои извинения!	 1
8	case_text1	Как нас найти?	 1
3	label_check-connection_window	Введите адрес:	Примечание в окне проверки подключения
4	button_check-connection_text	Проверить	Текст кнопки в окне проверки подключения
9	case_text2	Способы оплаты	 
11	section_phone_description	Оказание услуги реализовано на базе «Городской сети передачи данных» ООО «Орбител», построенной по технологии «оптика до дома» (FTTB). Подключение телефонного оборудования Абонентов к узлу ООО «Орбител» внутри здания осуществляется с помощью медного кабеля UTP, используемого для предоставления доступа к сети Интернет, что позволяет Абонентам пользоваться 2 услугами связи (интернет и телефон) независимо друг от друга, обеспечивая высокие качественные показатели этих услуг.	 
10	section_phone_title	Подключение телефонии от компании ООО "Орбител"	 
12	section_phone_card_title	Преимущества:	 
13	section_phone_card_elements	Тоновый режим набора номера;Бесплатный АОН;Низкая абонентская плата;Беспрецедентно низкая стоимость пользования 2-мя телефонными номерами;Закрепление телефонного номера за Абонентом при переездах (в пределах сети ООО «Орбител»).	 
14	section_phone_card_price	300 ₽/мес	 
16	section_inet-for-ul_title	Корпоративный интернет от ООО "Орбител"	 
15	section_phone_card_note	*требуется наличие VoIP-шлюза или IP-телефона	 
17	section_inet-for-ul_text	Подключение к сети Интернет от компании «Орбител» обеспечит вашему предприятию высокоскоростной и надежный доступ ко множеству ресурсов «Мировой паутины». Мы предлагаем высокое качество передачи данных на разных скоростях по оптимальным тарифным планам. Заключив договор на доступ к сети Интернет с нами, Вы сможете воспользоваться дополнительными сервисами, которые могут вам потребоваться в процессе работы.	 
18	sections_button-connect_text	Подключить	 
21	section_phone-for-ul_title	Телефон для бизнеса от ООО "Орбител"	 
22	section_phone-for-ul_text	Компания использует технологию подключения VoIP или поток E1, что позволяет использовать одну линию для подключения Интернета и Телефона	
25	section_vpn-for-ul_text1	Услуга по предоставлению защищенной общей сети для организации взаимодействия между территориально-распределенными объектами Клиента	 
19	sections_advantages_title	Наши преимущества	1
24	section_vpn-for-ul_title	Виртуальная частная сеть <span>VPN<span>	 1
26	section_vpn-for-ul_text2	Услуга <span>VPN<span> (Layer 3/2 VPN) состоит в предоставлении защищенной общей сети для организации связей между территориально-распределенными объектами Клиента с соблюдением параметров качества, определенных соглашением об уровне сервиса. Услуга предоставляется на базе магистральной сети «Орбител», использующей технологию MPLS.	 1
23	section_advantages_phone-for-ul_elements	Выгодные тарифные планы;Возможность объединения телефонных номеров в «серию»;Возможность выбора «красивых» номеров;В случае переезда – телефонный номер «переезжает» вмести с Вами	 1
33	section_routers_names	TP-Link 740/741/840/841;Tenda N300;D-Link Dir-300/615 Air-интерфейс;D-Link Dir-300/615 Silver-интерфейс	 1
29	section_virtual-ats_text	Если Вам необходимо быстро подключить телефон, организовать call-центр или удаленное рабочие место с минимальными вложениями, то облачная АТС от «Орбител» — лучшее решение поставленной задачи!	
28	section_virtual-ats_title	Виртуальная АТС	 
31	section_routers_title	Настройка роутера	 
32	label_routers_window	Выбирите свой роутер:	 
34	section_tv-manual_title	Настройка ТВ каналов	 
35	label_tv-manual_window	Выбирите свой телевизор:	 
36	section_oplata_title	Способы оплаты	 
38	section_calculator_title	Калькулятор тарифов	 
39	section_abonent-application_title	Оформление заявки	 
40	section_business-application_title	Оставить заявку на подключение	 
27	section_vpn-for-ul_text3	Создание защищенной корпоративной сети для эффективного обмена разнородной информацией;Объединение более двух удаленных офисов/объектов Клиента;Обеспечение высокой защищенности передачи информации;Интеграция всех видов трафика и данных в рамках одной услуги;Обеспечение высокой доступности сервисов;Оптимальный размер платежа за получаемые услуги	 1
30	section_advantage_virtual-ats_elements	Подключение данной услуги в любом месте, где есть Интернет.;Многоканальный номер – до Вас легко дозвонятся все клиенты и партнеры.;Голосовое меню;Конференц-связь;Интеграция с любой CRM;Запись и хранение разговоров	 1
\.


--
-- Data for Name: Tariff; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Tariff" (id, type_id, price, name, description, speed, digital_channel, analog_channel, image, color) FROM stdin;
32	1	650	Домашний-Ульта-2	Тариф рассчитан на большое количество устройств домашней сети 	300	157	54		#ee5037
33	2	400	домашний-Плюс	Тариф подходит для большого количества устройств 	100	0	0	speed-meter1.svg	#0177fd
34	2	500	Домашний-Ультра	Тариф подходит для большего количества устройств	300	0	0	speed-meter2.svg	#0177fd
37	3	200	Аналоговое ТВ	..	0	0	54		#0177fd
38	3	200	Цифровое ТВ 	 ... 	0	154	0		#0177fd
31	1	550	Домашний-Плюс-2	  Тариф рассчитан на несколько устройств домашней сети   	100	157	54		#0177fd
\.


--
-- Data for Name: Tariff_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Tariff_type" (id, name) FROM stdin;
1	Интернет+ТВ
2	Интернет
3	ТВ
\.


--
-- Data for Name: User; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."User" (id, name, phone, account_number, password, current_balance, current_tariff, role, house, flat, baned) FROM stdin;
67	Виноградова Людмила Игоревна 	9195641145	12345678	ffb3a44fab3851cbf5c51066aa60f2038aefab90d41f8e7a8eecd5e35965de7c	700	32	3	1	21	0
6	Иванов Иван Иванович	9195978629	77777777	ffb3a44fab3851cbf5c51066aa60f2038aefab90d41f8e7a8eecd5e35965de7c	1627.76	32	3	3	55	0
1	Admin	1	00000000	ffb3a44fab3851cbf5c51066aa60f2038aefab90d41f8e7a8eecd5e35965de7c	0	\N	1	1	21	0
\.


--
-- Name: Address_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Address_id_seq"', 10, true);


--
-- Name: Deposits_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Deposits_id_seq"', 11, true);


--
-- Name: Expenses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Expenses_id_seq"', 10, true);


--
-- Name: Faq_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Faq_id_seq"', 11, true);


--
-- Name: Role_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Role_id_seq"', 3, true);


--
-- Name: SEO_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."SEO_id_seq"', 21, true);


--
-- Name: Service_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Service_id_seq"', 23, true);


--
-- Name: Service_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Service_type_id_seq"', 4, true);


--
-- Name: Settings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Settings_id_seq"', 44, true);


--
-- Name: Tariff_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Tariff_id_seq"', 48, true);


--
-- Name: Type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Type_id_seq"', 3, true);


--
-- Name: User_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."User_id_seq"', 86, true);


--
-- Name: Deposit Deposits_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Deposit"
    ADD CONSTRAINT "Deposits_pk" PRIMARY KEY (id);


--
-- Name: Expense Expenses_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Expense"
    ADD CONSTRAINT "Expenses_pk" PRIMARY KEY (id);


--
-- Name: Role Role_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Role"
    ADD CONSTRAINT "Role_pk" PRIMARY KEY (id);


--
-- Name: SEO SEO_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."SEO"
    ADD CONSTRAINT "SEO_pk" PRIMARY KEY (id);


--
-- Name: Service Service_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Service"
    ADD CONSTRAINT "Service_pk" PRIMARY KEY (id);


--
-- Name: Service_type Service_type_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Service_type"
    ADD CONSTRAINT "Service_type_pk" PRIMARY KEY (id);


--
-- Name: Session Session_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Session"
    ADD CONSTRAINT "Session_pk" PRIMARY KEY (hash);


--
-- Name: Settings Settings_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Settings"
    ADD CONSTRAINT "Settings_pk" PRIMARY KEY (id);


--
-- Name: Tariff Tariff_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Tariff"
    ADD CONSTRAINT "Tariff_pk" PRIMARY KEY (id);


--
-- Name: Tariff_type Type_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Tariff_type"
    ADD CONSTRAINT "Type_pk" PRIMARY KEY (id);


--
-- Name: User User_account_number_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."User"
    ADD CONSTRAINT "User_account_number_key" UNIQUE (account_number);


--
-- Name: User User_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."User"
    ADD CONSTRAINT "User_pk" PRIMARY KEY (id);


--
-- Name: Address address_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Address"
    ADD CONSTRAINT address_pk PRIMARY KEY (id);


--
-- Name: Faq faq_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Faq"
    ADD CONSTRAINT faq_pk PRIMARY KEY (id);


--
-- Name: settings_key_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX settings_key_uindex ON public."Settings" USING btree (key);


--
-- Name: tariff_name_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX tariff_name_uindex ON public."Tariff" USING btree (name);


--
-- Name: Deposit Deposits_fk0; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Deposit"
    ADD CONSTRAINT "Deposits_fk0" FOREIGN KEY (user_id) REFERENCES public."User"(id);


--
-- Name: Expense Expenses_fk0; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Expense"
    ADD CONSTRAINT "Expenses_fk0" FOREIGN KEY (user_id) REFERENCES public."User"(id);


--
-- Name: Service Service_fk0; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Service"
    ADD CONSTRAINT "Service_fk0" FOREIGN KEY (type_id) REFERENCES public."Service_type"(id);


--
-- Name: Session Session_fk0; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Session"
    ADD CONSTRAINT "Session_fk0" FOREIGN KEY (user_id) REFERENCES public."User"(id);


--
-- Name: Tariff Tariff_fk0; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Tariff"
    ADD CONSTRAINT "Tariff_fk0" FOREIGN KEY (type_id) REFERENCES public."Tariff_type"(id);


--
-- Name: User User_fk0; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."User"
    ADD CONSTRAINT "User_fk0" FOREIGN KEY (current_tariff) REFERENCES public."Tariff"(id);


--
-- Name: User User_fk1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."User"
    ADD CONSTRAINT "User_fk1" FOREIGN KEY (role) REFERENCES public."Role"(id);


--
-- PostgreSQL database dump complete
--

