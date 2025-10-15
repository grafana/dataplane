# Table Kind

Status: EARLY Draft/Proposal

If a dataframe doesn't fit other contract, that will be considered as table kind contract. Following formats fall under the table kind contract.

- Table Plain
- Table Frame

## TablePlain Format

Version: 0.0

### Definition

This format is default contract for the table kind. This contract is mostly like SQL tables.

### Properties

- By default, it is single frame. If multiple frames present, all frame names have to be unique
- All field names must be unique to its frame. (No duplicate field names allowed within the frame)
- All the labels will be ignored

### Example

Following is an example of a table plain kind

| Name: User Name<br/>Type: []string<br/>Label: `nil` | Name: Country<br/>Type: []\*string<br/>Label: `nil` | Name: Age<br/>Type: []string<br/>Label: `nil` | Name: Is Employee<br/>Type: []boolean<br/>Label: `nil` |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------- | ------------------------------------------------------ |
| Foo                                                 | India                                               | 27                                            | false                                                  |
| Bar                                                 | USA                                                 | 35                                            | true                                                   |

## TableFrame Format

Version: 0.0

### Definition

This is the default fallback format for all the formats+kind.

### Properties

- There can be multiple frames
- Field name can be duplicate within frame. If duplicate field names found, they must be distinguished by labels
- Fields can have labels
