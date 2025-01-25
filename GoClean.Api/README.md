The API layer contains Gin handlers for RESTful APIs, a DI container and its configurations using the Dig framework, project-wide global variables, and database settings and migrations.<br>
Project documentation is prepared with swagger framework.<br>
This project supports multiple databases, including MSSQL, PostgreSQL, MongoDB, and Neo4j. Connection pooling management for each of these databases is handled using specific patterns developed in this layer.If these patterns are not followed, database connections increase disproportionately, 
leading to various database issues and potentially causing the database to crash.<br>
These patterns are developed using an approach where the database connection is first established in the main function of the project, followed by running periodic jobs to check the connection status and reconnect if needed.<br>
