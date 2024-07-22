BEGIN;
\c grow
DROP TABLE IF EXISTS diary, soil, week, fertilizer, harvest;
CREATE TABLE IF NOT EXISTS diary (
	id 				int,
	environment 	text,
	url 			text,
	seedbank 		text,
	strain 			text,
	isPhoto			boolean
);
CREATE TABLE IF NOT EXISTS soil (
	id 				int,
	sType			text,
	percentage 		int
);	
CREATE TABLE IF NOT EXISTS week (
	id 				int,
	week 			int,
	wType			text,
	height			double precision,
	tempDay			double precision,
	tempNight		double precision,
	humidity		double precision,
	potSize			double precision,
	water			double precision,
	ph 				double precision,
	lightS			int,
	tds				double precision,	
	--methods
	lst				boolean,
	hst 			boolean,
	sog				boolean,
	scrog			boolean,
	topping			boolean,
	fiming 			boolean,
	mainlining		boolean,
	defoliation		boolean,
	fromseed1212	boolean
);
CREATE TABLE IF NOT EXISTS fertilizer (
	id 				int,
	wId				int,
	name 			text,
	amount			double precision,
	href 			text
);
CREATE TABLE IF NOT EXISTS harvest (
	id 				int,
	wID				int,
	wetWeight		double precision,
	dryWeight		double precision,
	amountPlants	int,
	growRoomSize	double precision
);
copy diary FROM '/home/ramenbased/snek/gdscraper/data/output/diary.csv' WITH DELIMITER ',' NULL AS '';
copy soil FROM 'diary.csv' WITH DELIMITER ',' NULL AS '';
copy week FROM 'diary.csv' WITH DELIMITER ',' NULL AS '';
copy fertilizer FROM 'diary.csv' WITH DELIMITER ',' NULL AS '';
copy harvest FROM 'diary.csv' WITH DELIMITER ',' NULL AS '';
COMMIT;

