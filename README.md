# questions


``` postgresql

CREATE TABLE public.words (
id SERIAL PRIMARY KEY,
   word VARCHAR NOT NULL

);

CREATE TABLE public.room (
id SERIAL PRIMARY KEY,
  room bigint not null ,
  --  userid bigint not null,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  alive bool default true 

);

CREATE TABLE public.roomUser (
id SERIAL PRIMARY KEY,
  roomID bigint not null,
  userID bigint not null ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE public.roomwords (
id SERIAL PRIMARY KEY,
  roomID bigint not null,
  wordid bigint not null ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	
);

CREATE TABLE public.gameplay (
id SERIAL PRIMARY KEY,
  roomID bigint not null,
  playerid bigint not null ,
  message varchar(60), 
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	
);

```


```
ssh -R 80:localhost:3000 localhost.run
```