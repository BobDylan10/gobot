CREATE TABLE IF NOT EXISTS connectiontimes (
   id INT AUTO_INCREMENT PRIMARY KEY,
   player_id INT NOT NULL,
   c_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   minutes FLOAT NOT NULL,
   FOREIGN KEY (player_id)
      REFERENCES players(player_id)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS weaponskd (
   id INT AUTO_INCREMENT PRIMARY KEY,
   player_id INT NOT NULL,
   wid INT NOT NULL,
   deaths INT NOT NULL DEFAULT 0,
   kills INT NOT NULL DEFAULT 0,
   UNIQUE(player_id, wid),
   FOREIGN KEY (player_id)
      REFERENCES players(player_id)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS killson (
   id INT AUTO_INCREMENT PRIMARY KEY,
   player_id INT NOT NULL,
   victim_id INT NOT NULL,
   kills INT NOT NULL DEFAULT 0,
   UNIQUE(player_id, victim_id),
   FOREIGN KEY (player_id)
      REFERENCES players(player_id),
   FOREIGN KEY (victim_id)
      REFERENCES players(player_id)
) ENGINE=InnoDB;