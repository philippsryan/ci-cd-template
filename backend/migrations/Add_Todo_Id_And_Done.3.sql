ALTER TABLE Todo
ADD COLUMN Done BOOL,
ADD COLUMN DoneOn DATE,
ADD COLUMN Id INT AUTO_INCREMENT,
ADD PRIMARY KEY (Id);
