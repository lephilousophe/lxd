//go:build linux && cgo && !agent
// +build linux,cgo,!agent

package db

// The code below was generated by lxd-generate - DO NOT EDIT!

import (
	"database/sql"
	"fmt"

	"github.com/lxc/lxd/lxd/db/cluster"
	"github.com/lxc/lxd/lxd/db/query"
	"github.com/lxc/lxd/shared/api"
)

var _ = api.ServerEnvironment{}

var warningObjects = cluster.RegisterStmt(`
SELECT warnings.id, coalesce(nodes.name, '') AS node, coalesce(projects.name, '') AS project, coalesce(warnings.entity_type_code, -1), coalesce(warnings.entity_id, -1), warnings.uuid, warnings.type_code, warnings.status, warnings.first_seen_date, warnings.last_seen_date, warnings.updated_date, warnings.last_message, warnings.count
  FROM warnings LEFT JOIN nodes ON warnings.node_id = nodes.id LEFT JOIN projects ON warnings.project_id = projects.id
  ORDER BY warnings.uuid
`)

var warningObjectsByUUID = cluster.RegisterStmt(`
SELECT warnings.id, coalesce(nodes.name, '') AS node, coalesce(projects.name, '') AS project, coalesce(warnings.entity_type_code, -1), coalesce(warnings.entity_id, -1), warnings.uuid, warnings.type_code, warnings.status, warnings.first_seen_date, warnings.last_seen_date, warnings.updated_date, warnings.last_message, warnings.count
  FROM warnings LEFT JOIN nodes ON warnings.node_id = nodes.id LEFT JOIN projects ON warnings.project_id = projects.id
  WHERE warnings.uuid = ? ORDER BY warnings.uuid
`)

var warningObjectsByProject = cluster.RegisterStmt(`
SELECT warnings.id, coalesce(nodes.name, '') AS node, coalesce(projects.name, '') AS project, coalesce(warnings.entity_type_code, -1), coalesce(warnings.entity_id, -1), warnings.uuid, warnings.type_code, warnings.status, warnings.first_seen_date, warnings.last_seen_date, warnings.updated_date, warnings.last_message, warnings.count
  FROM warnings LEFT JOIN nodes ON warnings.node_id = nodes.id LEFT JOIN projects ON warnings.project_id = projects.id
  WHERE coalesce(project, '') = ? ORDER BY warnings.uuid
`)

var warningObjectsByStatus = cluster.RegisterStmt(`
SELECT warnings.id, coalesce(nodes.name, '') AS node, coalesce(projects.name, '') AS project, coalesce(warnings.entity_type_code, -1), coalesce(warnings.entity_id, -1), warnings.uuid, warnings.type_code, warnings.status, warnings.first_seen_date, warnings.last_seen_date, warnings.updated_date, warnings.last_message, warnings.count
  FROM warnings LEFT JOIN nodes ON warnings.node_id = nodes.id LEFT JOIN projects ON warnings.project_id = projects.id
  WHERE warnings.status = ? ORDER BY warnings.uuid
`)

var warningObjectsByNodeAndTypeCode = cluster.RegisterStmt(`
SELECT warnings.id, coalesce(nodes.name, '') AS node, coalesce(projects.name, '') AS project, coalesce(warnings.entity_type_code, -1), coalesce(warnings.entity_id, -1), warnings.uuid, warnings.type_code, warnings.status, warnings.first_seen_date, warnings.last_seen_date, warnings.updated_date, warnings.last_message, warnings.count
  FROM warnings LEFT JOIN nodes ON warnings.node_id = nodes.id LEFT JOIN projects ON warnings.project_id = projects.id
  WHERE coalesce(node, '') = ? AND warnings.type_code = ? ORDER BY warnings.uuid
`)

var warningObjectsByNodeAndTypeCodeAndProject = cluster.RegisterStmt(`
SELECT warnings.id, coalesce(nodes.name, '') AS node, coalesce(projects.name, '') AS project, coalesce(warnings.entity_type_code, -1), coalesce(warnings.entity_id, -1), warnings.uuid, warnings.type_code, warnings.status, warnings.first_seen_date, warnings.last_seen_date, warnings.updated_date, warnings.last_message, warnings.count
  FROM warnings LEFT JOIN nodes ON warnings.node_id = nodes.id LEFT JOIN projects ON warnings.project_id = projects.id
  WHERE coalesce(node, '') = ? AND warnings.type_code = ? AND coalesce(project, '') = ? ORDER BY warnings.uuid
`)

var warningObjectsByNodeAndTypeCodeAndProjectAndEntityTypeCodeAndEntityID = cluster.RegisterStmt(`
SELECT warnings.id, coalesce(nodes.name, '') AS node, coalesce(projects.name, '') AS project, coalesce(warnings.entity_type_code, -1), coalesce(warnings.entity_id, -1), warnings.uuid, warnings.type_code, warnings.status, warnings.first_seen_date, warnings.last_seen_date, warnings.updated_date, warnings.last_message, warnings.count
  FROM warnings LEFT JOIN nodes ON warnings.node_id = nodes.id LEFT JOIN projects ON warnings.project_id = projects.id
  WHERE coalesce(node, '') = ? AND warnings.type_code = ? AND coalesce(project, '') = ? AND coalesce(warnings.entity_type_code, -1) = ? AND coalesce(warnings.entity_id, -1) = ? ORDER BY warnings.uuid
`)

var warningDeleteByUUID = cluster.RegisterStmt(`
DELETE FROM warnings WHERE uuid = ?
`)

var warningDeleteByEntityTypeCodeAndEntityID = cluster.RegisterStmt(`
DELETE FROM warnings WHERE entity_type_code = ? AND entity_id = ?
`)

// GetWarnings returns all available warnings.
// generator: warning GetMany
func (c *ClusterTx) GetWarnings(filter WarningFilter) ([]Warning, error) {
	var err error

	// Result slice.
	objects := make([]Warning, 0)

	// Pick the prepared statement and arguments to use based on active criteria.
	var stmt *sql.Stmt
	var args []interface{}

	if filter.Node != nil && filter.TypeCode != nil && filter.Project != nil && filter.EntityTypeCode != nil && filter.EntityID != nil && filter.ID == nil && filter.UUID == nil && filter.Status == nil {
		stmt = c.stmt(warningObjectsByNodeAndTypeCodeAndProjectAndEntityTypeCodeAndEntityID)
		args = []interface{}{
			filter.Node,
			filter.TypeCode,
			filter.Project,
			filter.EntityTypeCode,
			filter.EntityID,
		}
	} else if filter.Node != nil && filter.TypeCode != nil && filter.Project != nil && filter.ID == nil && filter.UUID == nil && filter.EntityTypeCode == nil && filter.EntityID == nil && filter.Status == nil {
		stmt = c.stmt(warningObjectsByNodeAndTypeCodeAndProject)
		args = []interface{}{
			filter.Node,
			filter.TypeCode,
			filter.Project,
		}
	} else if filter.Node != nil && filter.TypeCode != nil && filter.ID == nil && filter.UUID == nil && filter.Project == nil && filter.EntityTypeCode == nil && filter.EntityID == nil && filter.Status == nil {
		stmt = c.stmt(warningObjectsByNodeAndTypeCode)
		args = []interface{}{
			filter.Node,
			filter.TypeCode,
		}
	} else if filter.UUID != nil && filter.ID == nil && filter.Project == nil && filter.Node == nil && filter.TypeCode == nil && filter.EntityTypeCode == nil && filter.EntityID == nil && filter.Status == nil {
		stmt = c.stmt(warningObjectsByUUID)
		args = []interface{}{
			filter.UUID,
		}
	} else if filter.Status != nil && filter.ID == nil && filter.UUID == nil && filter.Project == nil && filter.Node == nil && filter.TypeCode == nil && filter.EntityTypeCode == nil && filter.EntityID == nil {
		stmt = c.stmt(warningObjectsByStatus)
		args = []interface{}{
			filter.Status,
		}
	} else if filter.Project != nil && filter.ID == nil && filter.UUID == nil && filter.Node == nil && filter.TypeCode == nil && filter.EntityTypeCode == nil && filter.EntityID == nil && filter.Status == nil {
		stmt = c.stmt(warningObjectsByProject)
		args = []interface{}{
			filter.Project,
		}
	} else if filter.ID == nil && filter.UUID == nil && filter.Project == nil && filter.Node == nil && filter.TypeCode == nil && filter.EntityTypeCode == nil && filter.EntityID == nil && filter.Status == nil {
		stmt = c.stmt(warningObjects)
		args = []interface{}{}
	} else {
		return nil, fmt.Errorf("No statement exists for the given Filter")
	}

	// Dest function for scanning a row.
	dest := func(i int) []interface{} {
		objects = append(objects, Warning{})
		return []interface{}{
			&objects[i].ID,
			&objects[i].Node,
			&objects[i].Project,
			&objects[i].EntityTypeCode,
			&objects[i].EntityID,
			&objects[i].UUID,
			&objects[i].TypeCode,
			&objects[i].Status,
			&objects[i].FirstSeenDate,
			&objects[i].LastSeenDate,
			&objects[i].UpdatedDate,
			&objects[i].LastMessage,
			&objects[i].Count,
		}
	}

	// Select.
	err = query.SelectObjects(stmt, dest, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"warnings\" table: %w", err)
	}

	return objects, nil
}

// GetWarning returns the warning with the given key.
// generator: warning GetOne-by-UUID
func (c *ClusterTx) GetWarning(uuid string) (*Warning, error) {
	filter := WarningFilter{}
	filter.UUID = &uuid

	objects, err := c.GetWarnings(filter)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"warnings\" table: %w", err)
	}

	switch len(objects) {
	case 0:
		return nil, ErrNoSuchObject
	case 1:
		return &objects[0], nil
	default:
		return nil, fmt.Errorf("More than one \"warnings\" entry matches")
	}
}

// DeleteWarning deletes the warning matching the given key parameters.
// generator: warning DeleteOne-by-UUID
func (c *ClusterTx) DeleteWarning(uuid string) error {
	stmt := c.stmt(warningDeleteByUUID)
	result, err := stmt.Exec(uuid)
	if err != nil {
		return fmt.Errorf("Delete \"warnings\": %w", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Fetch affected rows: %w", err)
	}

	if n != 1 {
		return fmt.Errorf("Query deleted %d rows instead of 1", n)
	}

	return nil
}

// DeleteWarnings deletes the warning matching the given key parameters.
// generator: warning DeleteMany-by-EntityTypeCode-and-EntityID
func (c *ClusterTx) DeleteWarnings(entityTypeCode int, entityID int) error {
	stmt := c.stmt(warningDeleteByEntityTypeCodeAndEntityID)
	result, err := stmt.Exec(entityTypeCode, entityID)
	if err != nil {
		return fmt.Errorf("Delete \"warnings\": %w", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Fetch affected rows: %w", err)
	}

	return nil
}
