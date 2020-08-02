### Todo

 - [x] ProtoBuf
 - [x] Add Service
 - [x] Implement Service
 - [x] Add Repository
 - [x] Implement Repository
 - [x] Add GRPC Server
 - [x] Add GRPC Client
 - [x] Postgres Migrations
 - [x] Add main.go on Each Service
 - [x] Fix Postgres Lib
 - [x] GraphQL Schema
 - [x] GraphQL Query Resolver
 - [ ] GraphQL Mutation Resolver
 - [x] Fix GetByID Client
 - [ ] docker-compose.yaml
 - [ ] Auth Service
 - [ ] Deploy
 

### Database Schema

1. Role
   - ID
   - Name
   - Description
   - CreatedAt
   - UpdatedAt

2. Gender
   - ID
   - Name
   - Description
   - CreatedAt
   - UpdatedAt

3. Blood
   - ID
   - Name
   - Rhesus
   - Description
   - CreatedAt
   - UpdatedAt
   
4. History
   - User
   - Disease
   - Note


### MongoDB

1. User
   - ID
   - Name
   - Email
   - Address
   - RoleID
   - GenderID
   - BloodID
   - BirthOfDate
   - Contact

### Elastic Search

1. Disease
   - ID
   - Name
   - Description
   - CreatedAt
   - UpdatedAt
