CREATE TABLE public."Product" (
    id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    price_in_cents integer NOT NULL,
    title text NOT NULL,
    description text NOT NULL
);
COMMENT ON TABLE public."Product" IS 'Describe all products that are on sale.';
CREATE TABLE public."User" (
    id text NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,
    date_of_birth date NOT NULL
);
COMMENT ON TABLE public."User" IS 'Describe users, with basic data.';
ALTER TABLE ONLY public."Product"
    ADD CONSTRAINT "Product_pkey" PRIMARY KEY (id, title);
ALTER TABLE ONLY public."Product"
    ADD CONSTRAINT "Product_title_key" UNIQUE (title);
ALTER TABLE ONLY public."User"
    ADD CONSTRAINT "User_id_key" UNIQUE (id);
ALTER TABLE ONLY public."User"
    ADD CONSTRAINT "User_pkey" PRIMARY KEY (id);
