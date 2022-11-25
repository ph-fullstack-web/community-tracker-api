ALTER TABLE public.project 
	ALTER COLUMN projectcode TYPE varchar(100),
	ALTER COLUMN projectcode SET NOT null,
	ALTER COLUMN projectcode SET DEFAULT 'PROJECT-CODE'
