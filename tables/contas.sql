CREATE TABLE contas ( 
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY, 
    agencia numeric(4,0) NOT NULL, 
    conta numeric(10,0) NOT NULL, 
    senha varchar(255) NOT NULL 
);