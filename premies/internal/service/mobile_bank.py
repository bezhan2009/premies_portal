import pandas as pd

from internal.repository import mobile_bank


def mobile_bank_excel_upload(path_file: str) -> Exception | str:
    df = pd.read_excel(path_file)
    df.columns = df.columns.str.strip().str.lower()  # нормализация названий столбцов
    df.rename(columns={'surname': 'surname', 'inn': 'inn'}, inplace=True)
    df['inn'] = df['inn'].astype(str)

    mobile_bank_clean_table()

    return mobile_bank.mobile_bank_excel_upload(df)


def mobile_bank_clean_table() -> Exception | str:
    return mobile_bank.mobile_bank_clean_table()
