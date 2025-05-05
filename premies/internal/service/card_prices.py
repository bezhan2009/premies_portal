import pandas as pd

from internal.repository import card_prices

coast_dict = {}


def upload_card_prices(path_file: str):
    # Загружаем прайс-лист стоимости карт
    price_df = pd.read_excel(path_file, engine='openpyxl')

    return card_prices.upload_card_prices(price_df)


def upload_card_prices_to_dict():
    global coast_dict

    res_dict = card_prices.upload_card_prices_to_dict()
    if type(res_dict) is Exception:
        return res_dict

    coast_dict = res_dict
    return "Successfully uploaded card prices to dictionary."


def get_coast_dict() -> dict:
    global coast_dict
    return coast_dict
