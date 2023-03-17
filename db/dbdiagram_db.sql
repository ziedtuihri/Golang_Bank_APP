// dbdiagram.io

Table accounts as A{
  id bigserial [pk]
  owner varchar [not null]
  balance bigint [not null]
  currency varchar [not null]
  created_at timestamptz [default: `now()`]

  Indexes {
    id [pk]
  }
}

Table entries as E {
  id bigserial [pk]
  account_id bigint [ref: > A.id, not null]
  amount bigint [not null]
  created_at timestampts [default: `now()`]
}

Table transfers {
  id bigserial [pk]
  from_account_id bidint [ref: > A.id, not null]
  to_account_id bigint [ref: > A.id, not null]
  amount bigint [not null, note: "must be positive"]
  created_at timestamptz [default: `now()`]

    Indexes {
    from_account_id
    to_account_id
    (from_account_id, to_account_id)
  }
}




