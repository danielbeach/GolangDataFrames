import pandas as pd
from datetime import datetime

def main():
    t1 = datetime.now()
    df = pd.read_csv("data/202206-divvy-tripdata.csv")
    df = df[df.member_casual == 'member']
    df2 = df.groupby(['start_station_name'])['ride_id'].count().reset_index(name='count') \
                             .sort_values(['count'], ascending=False)
    print(df2)

    t2 = datetime.now()
    print("it took {x} to run".format(x=t2-t1))

if __name__ == '__main__':
    main()
