import pandas as pd

from internal.repository import cards
from internal.service.card_prices import get_coast_dict
from pkg.logger.logger import setup_logger

logger = setup_logger(__name__)


def upload_cards(file_path: str) -> Exception | str:
    OP = "service.upload_cards"

    try:
        df = pd.read_excel(file_path, engine='openpyxl')
    except FileNotFoundError as e:
        logger.error("[{}] File not found {}".format(OP, file_path))
        return e

    clean_cards_table()

    return cards.upload_cards(df, get_coast_dict())


def clean_cards_table() -> Exception | str:
    return cards.clean_cards_table()
