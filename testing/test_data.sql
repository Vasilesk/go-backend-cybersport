insert into games(name, description) values('Dota 2', 'Dota 2 is very popular game I`ve never played');
insert into games(name, description) values('Warcraft 2', 'Warcraft 2 not very popular game but I`ve played it');
insert into players(name, description, rating) values('Petya', 'Petya is a player with the highest rating', 1);
insert into players(name, description, rating) values('Kolya', 'Kolya is not a good player. He is a younger brother of Petya', 0.1);
insert into teams(name, description, game_id) values('Petya Inc', 'Petya`s team', 1);
insert into teams_players(team_id, player_id, in_team_player_id) values(1, 1, 1);
insert into teams_players(team_id, player_id, in_team_player_id) values(1, 2, 2);
