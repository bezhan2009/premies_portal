import pandas as pd

from internal.repository import tus


def tus_excel_upload(file_path: str) -> Exception | str:
    df = pd.read_excel(file_path)
    df.columns = df.columns.str.strip().str.lower()

    df.rename(columns={
        'dvid': 'dvid',
        'reqdt': 'req_date',
        'code': 'code',
        'tus_code': 'tus_code',
        'mark': 'mark'
    }, inplace=True)

    df['dvid'] = pd.to_datetime(df['dvid'], dayfirst=True, errors='coerce')
    df['req_date'] = pd.to_datetime(df['req_date'], dayfirst=True, errors='coerce')

    tus_clean_table()

    return tus.tus_excel_upload(df)


def tus_clean_table() -> Exception | str:
    return tus.tus_clean_table()
