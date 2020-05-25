CREATE TABLE usuario(
ID_USUAR INT AUTO_INCREMENT PRIMARY KEY,
NOME_USUAR VARCHAR(30) NOT NULL,
APELIDO_USUAR VARCHAR(15) NOT NULL,
EMAIL_USUAR VARCHAR(30) NOT NULL,
AVATAR_USUAR VARCHAR(254) NOT NULL
);

CREATE TABLE mesa(
ID_MESA INT AUTO_INCREMENT PRIMARY KEY,
ADM_MESA VARCHAR(15),
TITULO_MESA VARCHAR (20) NOT NULL,
DESC_MESA VARCHAR (100) NOT NULL,
QTDEJOG_MESA INT NOT NULL,
FORMA_MESA VARCHAR (50) NOT NULL,
STATUS_MESA VARCHAR (15) NOT NULL,
DATCRIAC_MESA TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE jogadores(
ID_JOGA INT AUTO_INCREMENT PRIMARY KEY,
ID_MESA INT NOT NULL,
ID_USUAR INT NOT NULL,
MESTRE_JOGA BOOL NOT NULL,
FICHA_JOGA VARCHAR(254),
FOREIGN KEY (ID_USUAR) REFERENCES usuario(ID_USUAR),
FOREIGN KEY (ID_MESA) REFERENCES mesa(ID_MESA)
);

CREATE TABLE sessoes_ativas(
ID_SESS INT AUTO_INCREMENT PRIMARY KEY,
ID_USUAR INT,
DATA_SESS TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
FOREIGN KEY (ID_USUAR) REFERENCES usuario(ID_USUAR)
)