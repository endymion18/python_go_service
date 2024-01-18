import datetime
import json

from asyncpg import Connection

from greeting.exceptions import LengthException


async def check_name(name: str) -> None:
    if len(name) == 0 or name.isspace():
        raise LengthException("Name length can't be null")
    if len(name) > 30:
        raise LengthException("Name length is too long")


async def add_greeting_to_db(name: str, conn: Connection):
    timestamp = datetime.datetime.now()
    stmt = f"""INSERT INTO greet_info VALUES
    (default, '{name}', '{timestamp}');
    """
    await conn.execute(stmt)


async def get_all_greetings_from_db(conn: Connection):
    stmt = f"""SELECT * FROM greet_info"""
    rows = await conn.fetch(stmt)

    return rows
