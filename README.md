# Patterns of Enterprise Application Architecture

[![main](https://github.com/flowck/patterns_of_enterprise_application_architecture_golang/actions/workflows/main.yml/badge.svg)](https://github.com/flowck/patterns_of_enterprise_application_architecture_golang/actions/workflows/main.yml)

This repository has been created with the intent of demonstrate the implementation of some patterns described in
the book "Patterns of Enterprise Application Architecture" by Martin Fowler, D. Rice, M. Foemmel, E. Hieatt, R. Mee, and
R. Stafford.

<div style="display: flex; justify-content: center;">
  <img src="cover.jpg" width="512" />
</div>

## Patterns Implemented

- [Row Data Gateway](row_data_gateway): An object that acts as a Gateway (466) to a single record in a data source. There is one instance per row. [[1]](#ref-1)
- [Optimistic Offline Lock](optimistic_offline_lock): Prevents conflicts between concurrent business process transactions by detecting a conflict and rolling back the transaction. [[1]](#ref-1)

## References

<ul>
  <li id="ref-1">[1] Patterns of Enterprise Application Architecture" by Martin Fowler, D. Rice, M. Foemmel, E. Hieatt, R. Mee, and
  R. Stafford.</li>
</ul>
