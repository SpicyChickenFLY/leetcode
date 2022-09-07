select
  distinct a.Id
from
  Weather as a
  Join Weather as b on datediff(b.RecordDate, a.RecordDate) = -1
where
  a.Temperature > b.Temperature;