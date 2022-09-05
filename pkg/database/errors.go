package database

import (
    "fmt"
)

type ErrOperation struct {
    operation    string
}

func (e *ErrOperation) Error() string {
    return fmt.Sprintf("couldn't perform operation: %s", e.operation)
}

type ErrConnectDatabase struct{}

func (e *ErrConnectDatabase) Error() string {
    return fmt.Sprintf("couldn't connect to database")
}

type ErrCreateDatabase struct{}

func (e *ErrCreateDatabase) Error() string {
    return fmt.Sprintf("couldn't create database")
}

type ErrUnimplementedDatabase struct {
    database    string
}

func (e *ErrUnimplementedDatabase) Error() string {
    return fmt.Sprintf("database %s not implemented", e.database)
}

