CREATE TABLE Todo (
	Title varchar(255),
	Body varchar(255),
	BelongsTo varchar(255),
	FOREIGN KEY (BelongsTo) REFERENCES User(Name)
);
