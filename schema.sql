CREATE TABLE informacion_admin (
	codigo_admin bigint,
	nombre varchar(255),
	apellido varchar(255)
);
ALTER TABLE informacion_admin ADD CONSTRAINT codigo_admin PRIMARY KEY(codigo_admin);
CREATE TABLE login_admin (
	password varchar(255),
	FK_informacion_admin_codigo_admin bigint
);
CREATE TABLE informacion_estudiante (
	codigo_estudiante bigint,
	nombre varchar(255),
	apellido varchar(255)
);
ALTER TABLE informacion_estudiante ADD CONSTRAINT codigo_estudiante PRIMARY KEY(codigo_estudiante);
CREATE TABLE login_estudiante (
	password varchar(255),
	FK_informacion_estudiante_codigo_estudiante bigint
);
CREATE TABLE inscripcion_comedor (
	contador serial,
	FK_informacion_estudiante_codigo_estudiante bigint
);
CREATE TABLE info_qr (
	qr text
);
ALTER TABLE login_admin ADD CONSTRAINT codigo_admin_login FOREIGN KEY (FK_informacion_admin_codigo_admin) REFERENCES informacion_admin(codigo_admin) ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE login_estudiante ADD CONSTRAINT codigo_estudiante_login FOREIGN KEY (FK_informacion_estudiante_codigo_estudiante) REFERENCES informacion_estudiante(codigo_estudiante) ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE inscripcion_comedor ADD CONSTRAINT inscripcion FOREIGN KEY (FK_informacion_estudiante_codigo_estudiante) REFERENCES informacion_estudiante(codigo_estudiante) ON DELETE NO ACTION ON UPDATE NO ACTION;
