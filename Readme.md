# Inventory

An app I re-write every time I want to learn a new web language.

This time with golang.

## Things to do

Inventory application for media.

### Models

#### Item

- Name
- Year
- Notes
- Creator_id
- Format_id

#### Creator

- Name
- Notes

#### Format

- Name

### Features

- [ ] Index page
- [ ] List of all media
- [ ] List of media by Format
- [ ] List of media by Creator
- [ ] Add items
- [ ] View item
- [ ] Edit item
- [ ] Authentication


#### gorilla/context

This is for global variables passed between middlewares when using normal `http.HandleFunc` type methods

http://www.gorillatoolkit.org/pkg/context