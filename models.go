package main

import (
  "time"
)

type Thing struct {
  Id          int64
  AuthorId    int64
  PublisherId int64
  Year        int64
  CreatedAt   time.Time
  UpdatedAt   time.Time
}

type Author struct {
  Id            int64
  CreatedAt     time.Time
  UpdatedAt     time.Time
  Name          string
}

type Publisher struct {
  Id            int64
  CreatedAt     time.Time
  UpdatedAt     time.Time
  Name          string
}