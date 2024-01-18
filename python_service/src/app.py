from fastapi import FastAPI
from greeting.router import greeting_router

app = FastAPI(
    title="Privet API"
)

app.include_router(greeting_router)

