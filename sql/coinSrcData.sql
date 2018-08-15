
use coin_data;

-- 1min, 5min, 15min, 30min, 60min, 1day, 1mon, 1week, 1year
CREATE TABLE `tb_coinData_1min` (
  `coinType` tinyint(3) unsigned NOT NULL COMMENT,
  `ID` bigint(20) unsigned NOT NULL,
  `Amount` decimal(20,10) DEFAULT NULL,
  `Count` bigint(20) DEFAULT NULL,
  `Open` decimal(20,10) DEFAULT NULL,
  `Close` decimal(20,10) DEFAULT NULL,
  `Low` decimal(20,10) DEFAULT NULL,
  `High` decimal(20,10) DEFAULT NULL,
  `Vol` decimal(20,10) DEFAULT NULL,
  PRIMARY KEY (`coinType`,`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tb_coinData_5min` (
  `coinType` tinyint(3) unsigned NOT NULL COMMENT,
  `ID` bigint(20) unsigned NOT NULL,
  `Amount` decimal(20,10) DEFAULT NULL,
  `Count` bigint(20) DEFAULT NULL,
  `Open` decimal(20,10) DEFAULT NULL,
  `Close` decimal(20,10) DEFAULT NULL,
  `Low` decimal(20,10) DEFAULT NULL,
  `High` decimal(20,10) DEFAULT NULL,
  `Vol` decimal(20,10) DEFAULT NULL,
  PRIMARY KEY (`coinType`,`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tb_coinData_15min` (
  `coinType` tinyint(3) unsigned NOT NULL COMMENT,
  `ID` bigint(20) unsigned NOT NULL,
  `Amount` decimal(20,10) DEFAULT NULL,
  `Count` bigint(20) DEFAULT NULL,
  `Open` decimal(20,10) DEFAULT NULL,
  `Close` decimal(20,10) DEFAULT NULL,
  `Low` decimal(20,10) DEFAULT NULL,
  `High` decimal(20,10) DEFAULT NULL,
  `Vol` decimal(20,10) DEFAULT NULL,
  PRIMARY KEY (`coinType`,`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tb_coinData_30min` (
  `coinType` tinyint(3) unsigned NOT NULL COMMENT,
  `ID` bigint(20) unsigned NOT NULL,
  `Amount` decimal(20,10) DEFAULT NULL,
  `Count` bigint(20) DEFAULT NULL,
  `Open` decimal(20,10) DEFAULT NULL,
  `Close` decimal(20,10) DEFAULT NULL,
  `Low` decimal(20,10) DEFAULT NULL,
  `High` decimal(20,10) DEFAULT NULL,
  `Vol` decimal(20,10) DEFAULT NULL,
  PRIMARY KEY (`coinType`,`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tb_coinData_60min` (
  `coinType` tinyint(3) unsigned NOT NULL COMMENT,
  `ID` bigint(20) unsigned NOT NULL,
  `Amount` decimal(20,10) DEFAULT NULL,
  `Count` bigint(20) DEFAULT NULL,
  `Open` decimal(20,10) DEFAULT NULL,
  `Close` decimal(20,10) DEFAULT NULL,
  `Low` decimal(20,10) DEFAULT NULL,
  `High` decimal(20,10) DEFAULT NULL,
  `Vol` decimal(20,10) DEFAULT NULL,
  PRIMARY KEY (`coinType`,`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tb_coinData_1day` (
  `coinType` tinyint(3) unsigned NOT NULL COMMENT,
  `ID` bigint(20) unsigned NOT NULL,
  `Amount` decimal(20,10) DEFAULT NULL,
  `Count` bigint(20) DEFAULT NULL,
  `Open` decimal(20,10) DEFAULT NULL,
  `Close` decimal(20,10) DEFAULT NULL,
  `Low` decimal(20,10) DEFAULT NULL,
  `High` decimal(20,10) DEFAULT NULL,
  `Vol` decimal(20,10) DEFAULT NULL,
  PRIMARY KEY (`coinType`,`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;






