from asyncpg import Connection

from greeting.exceptions import LengthException


async def check_name(name: str) -> None:
    """Check if name is correct
    :param name: Name to check
    :return: None or raises LengthException
    """
    if len(name) == 0 or name.isspace():
        raise LengthException("Name length can't be null")
    if len(name) > 30:
        raise LengthException("Name length is too long")


async def add_greeting_to_db(name: str, conn: Connection) -> None:
    """Check if name is correct
    :param name: Name to insert into table
    :param conn: Database connection
    """
    stmt = f"""INSERT INTO greet_info_python VALUES
    (default, '{name}', default);
    """
    await conn.execute(stmt)


async def get_all_greetings_from_db(conn: Connection):
    """
    Get all rows from greeting table
    :param conn: Database connection
    :return: All rows from table
    """
    stmt = f"""SELECT * FROM greet_info_python"""
    rows = await conn.fetch(stmt)

    return rows
