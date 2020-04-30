-- +goose Up
-- SQL in this section is executed when the migration is applied.

INSERT INTO livers (id, name, channel_id, main_color, sub_color) VALUES
  ('tokino-sora',      'ときのそら',           'UCp6993wxpyDPHUpavwDFqgg', '#4374FF', '#FD00AE'),
  ('roboco-san',       'ロボ子さん',           'UCDqI2jOz0weumE8s7paEk6g', '#D252FF', '#FEA5D1'),
  ('sakuramiko',       'さくらみこ',           'UC-hM6YJuNYVAmUWxeIr9FeA', '#FEA5D1', '#F7002F'),
  ('yozora-mel',       '夜空メル',             'UCD8HOxPs4Xvsm8H0ZxXGiBw', '#FFD200', '#FF7608'),
  ('shirakami-fubuki', '白上フブキ',           'UCdn5BQ06XqgXoAxIhbqw5Rg', '#49E5FF', '#FFFFFF'),
  ('natsuiro-matsuri', '夏色まつり',           'UCQ0UDLQCjY0rmuxCDE38FGg', '#FF7608', '#3BE898'),
  ('akai-haato',       '赤井はあと',           'UC1CfXB_kRs3C-zaeTG3oGyg', '#FD00AE', '#F7002F'),
  ('aki-rosenthal',    'アキ・ローゼンタール', 'UCFTLzh12_nrtzqBPsTCqenA', '#8045FF', '#FEA5D1'),
  ('minato-aqua',      '湊あくあ',             'UC1opHUrw8rvnsadT-iGp7Cg', '#D252FF', '#49E5FF'),
  ('nakiri-ayame',     '百鬼あやめ',           'UC7fk0CB07ly8oSl0aqKkqFg', '#F7002F', '#FFFFFF'),
  ('yuzuki-choco',     '癒月ちょこ',           'UC1suqwovbL1kzsoaZgFZLKg', '#FD00AE', '#FEA5D1'),
  ('murasaki-shion',   '紫咲シオン',           'UCXTpFs_3PqI41qX2d9tL2Rw', '#D252FF', '#FFD200'),
  ('oozora-subaru',    '大空スバル',           'UCvzGlP9oQwU--Y0r9id_jnA', '#FFD200', '#3BE898'),
  ('ookami-mio',       '大神ミオ',             'UCp-5t9SrOQwXMU7iIjQfARg', '#3BE898', '#F7002F'),
  ('nekomata-okayu',   '猫又おかゆ',           'UCvaTdHTWBGv3MKj3KVqJVCw', '#D252FF', '#4374FF'),
  ('inugami-korone',   '戌神ころね',           'UChAnqc_AY5_I3Px5dig3X1Q', '#FFD200', '#49E5FF'),
  ('shiranui-flare',   '不知火フレア',         'UCvInZx9h3jC2JzsIzoOebWg', '#FF7608', '#FFFFFF'),
  ('shirogane-noel',   '白銀ノエル',           'UCdyqAaZDKHXg4Ahi7VENThQ', '#FFFFFF', '#49E5FF'),
  ('houshou-marine',   '宝鐘マリン',           'UCCzUftO8KOVkV4wQG1vkUvg', '#F7002F', '#FF7608'),
  ('usada-pekora',     '兎田ぺこら',           'UC1DCedRgGHBdm81E1llLhOQ', '#49E5FF', '#FF7608'),
  ('uruha-rushia',     '潤羽るしあ',           'UCl_gCybOJRIgOXw6Qb4qJzQ', '#8045FF', '#93F340'),
  ('hoshimati-suisei', '星街すいせい',         'UC5CwaMl1eIgY8h02uZw7u8A', '#4374FF', '#49E5FF'),
  ('amane-kanata',     '天音かなた',           'UCZlDXzGoo7d44bwdNObFacg', '#FFFFFF', '#FFD200'), -- FIXME: color
  ('kiryu-coco',       '桐生ココ',             'UCS9uQI-jC3DE0L4IpXyvr6w', '#FF7608', '#F7002F'), -- FIXME: color
  ('tsunomaki-watame', '角巻わため',           'UCqm3BQLlJfvkTsX_hvm0UmA', '#FFD200', '#FFFFFF'), -- FIXME: color
  ('tokoyami-towa',    '常闇トワ',             'UC1uv2Oq6kNxgATlCiez59hw', '#D252FF', '#FEA5D1'), -- FIXME: color
  ('himemori-luna',    '姫森ルーナ',           'UCa9Y57gfeY0Zro_noHRVrnw', '#8045FF', '#FD00AE'); -- FIXME: color

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

TRUNCATE TABLE livers;

