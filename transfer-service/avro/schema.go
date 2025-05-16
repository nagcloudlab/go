package avro

const TransferEventSchema = `
{
  "type": "record",
  "name": "TransferEvent",
  "fields": [
    {"name": "from_account", "type": "string"},
    {"name": "to_account", "type": "string"},
    {"name": "amount", "type": "double"}
  ]
}
`
