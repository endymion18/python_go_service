from asyncpg import Connection
from fastapi import APIRouter, Depends
from starlette.responses import JSONResponse

from greeting.exceptions import LengthException
from greeting.utils import check_name, add_greeting_to_db, get_all_greetings_from_db
from src.database import get_connection

greeting_router = APIRouter()


@greeting_router.get("/greet")
async def greet_by_name(name: str, connection: Connection = Depends(get_connection)) -> JSONResponse:
    """
    Takes name, returns JSON-object

    :param name: Name to greet
    :param connection: Get database connection using dependency
    :return: JSONResponse
    """
    try:
        await check_name(name)
    except LengthException as ex:
        return JSONResponse(content=ex.__str__(), status_code=400)

    await add_greeting_to_db(name, connection)
    return JSONResponse(content=f"Привет, {name} от Python!", status_code=200)


@greeting_router.get("/greet/history")
async def get_all_greetings(connection: Connection = Depends(get_connection)) -> JSONResponse:
    """
    :param connection: Get database connection using dependency
    :return: JSON list of all table rows
    """
    all_greetings = await get_all_greetings_from_db(connection)

    return all_greetings
