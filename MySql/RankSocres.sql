# Write your MySQL query statement below
select Scores.Score, Rank from Scores LEFT JOIN (
    select NewScore.Score, (@i:=@i+1) AS Rank from (
        SELECT DISTINCT Score FROM Scores ORDER BY Score DESC
    ) AS NewScore, (SELECT @i:=0) AS It
) AS Temp ON Scores.Score = Temp.Score ORDER BY Rank;