CREATE TABLE games (
    id serial PRIMARY KEY,
    name text,
    description text,
    logo_link text
);

CREATE TABLE players (
    id serial PRIMARY KEY,
    name text,
    description text,
    logo_link text,
    rating real
);

CREATE TABLE teams (
    id serial PRIMARY KEY,
    name text,
    description text,
    logo_link text,
    rating real,
    game_id integer REFERENCES games (id) ON DELETE CASCADE
);

CREATE TABLE teams_players (
    -- id serial PRIMARY KEY,
    team_id integer REFERENCES teams (id) ON DELETE CASCADE,
    player_id integer REFERENCES players (id) ON DELETE CASCADE,
    in_team_player_id integer
);

CREATE TABLE tournaments (
    id serial PRIMARY KEY,
    name text,
    description text,
    logo_link text,
    is_active boolean,
    game_id integer REFERENCES games (id) ON DELETE CASCADE
);

CREATE TABLE tournaments_teams (
    tournament_id integer REFERENCES tournaments (id) ON DELETE CASCADE,
    team_id integer REFERENCES teams (id) ON DELETE CASCADE,
    in_tournament_team_id integer,
    rating real
);

CREATE TABLE matches (
    id serial PRIMARY KEY,
    in_tournament_match_id integer,
    tournament_id integer REFERENCES tournaments (id) ON DELETE CASCADE,
    name text,
    description text,
    logo_link text,
    video_link text,
    result text,
    m_date date
);

CREATE TABLE matches_teams (
    match_id integer REFERENCES matches (id) ON DELETE CASCADE,
    team_id integer REFERENCES teams (id) ON DELETE CASCADE,
    in_match_team_id integer UNIQUE
);

CREATE TABLE in_matches_players (
    match_id integer REFERENCES matches (id) ON DELETE CASCADE,
    in_match_team_id integer REFERENCES matches_teams (in_match_team_id) ON DELETE CASCADE,
    player_id integer REFERENCES players (id) ON DELETE CASCADE
);
