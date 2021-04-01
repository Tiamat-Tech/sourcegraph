package lsifstore

import (
	"database/sql"

	"github.com/sourcegraph/sourcegraph/enterprise/lib/codeintel/semantic"
	"github.com/sourcegraph/sourcegraph/internal/database/basestore"
)

type QualifiedDocumentData struct {
	UploadID int
	semantic.KeyedDocumentData
}

type DocumentDataColumn uint8

const (
	DocumentDataColumnRanges DocumentDataColumn = 1 << iota
	DocumentDataColumnHoverText
	DocumentDataColumnMonikers
	DocumentDataColumnPackageInformation
	DocumentDataColumnDiagnostics

	DocumentDataColumnAll = DocumentDataColumnRanges | DocumentDataColumnHoverText | DocumentDataColumnMonikers | DocumentDataColumnPackageInformation | DocumentDataColumnDiagnostics
)

// makeDocumentDataScanner creates a scanner that reads qualified document data from its given
// row object. Only the values indicated by the columns bitset will be populated in the document
// payload structs.
func (s *Store) makeDocumentDataScanner(columns DocumentDataColumn) func(rows *sql.Rows, queryErr error) (_ []QualifiedDocumentData, err error) {
	return func(rows *sql.Rows, queryErr error) (_ []QualifiedDocumentData, err error) {
		if queryErr != nil {
			return nil, queryErr
		}
		defer func() { err = basestore.CloseRows(rows, err) }()

		var values []QualifiedDocumentData
		for rows.Next() {
			record, err := s.scanSingleDocumentDataObject(rows, DocumentDataColumnAll)
			if err != nil {
				return nil, err
			}

			values = append(values, record)
		}

		return values, nil
	}
}

// makeDocumentVisitor returns a function that calls the given visitor function over matching
// document values. Only the values indicated by the columns bitset will be populated in the
// document payload structs.
func (s *Store) makeDocumentVisitor(columns DocumentDataColumn, f func(string, semantic.DocumentData)) func(rows *sql.Rows, queryErr error) error {
	return func(rows *sql.Rows, queryErr error) (err error) {
		if queryErr != nil {
			return queryErr
		}
		defer func() { err = basestore.CloseRows(rows, err) }()

		for rows.Next() {
			record, err := s.scanSingleDocumentDataObject(rows, DocumentDataColumnAll)
			if err != nil {
				return err
			}

			f(record.Path, record.Document)
		}

		return nil
	}
}

// makeFirstDocumentDataScanner creates a scanner that reads qualified document data from its given
// row object and returns the first one. If no rows match the query, a false-valued flag is returned.
// Only the values indicated by the columns bitset will be populated in the document payload structs.
func (s *Store) makeFirstDocumentDataScanner(columns DocumentDataColumn) func(rows *sql.Rows, queryErr error) (_ QualifiedDocumentData, _ bool, err error) {
	return func(rows *sql.Rows, queryErr error) (_ QualifiedDocumentData, _ bool, err error) {
		if queryErr != nil {
			return QualifiedDocumentData{}, false, queryErr
		}
		defer func() { err = basestore.CloseRows(rows, err) }()

		if rows.Next() {
			record, err := s.scanSingleDocumentDataObject(rows, DocumentDataColumnAll)
			if err != nil {
				return QualifiedDocumentData{}, false, err
			}

			return record, true, nil
		}

		return QualifiedDocumentData{}, false, nil
	}
}

// scanSingleDocumentDataObject populates a qualified document data value from the
// given cursor.
func (s *Store) scanSingleDocumentDataObject(rows *sql.Rows, columns DocumentDataColumn) (QualifiedDocumentData, error) {
	var rawData []byte
	var record QualifiedDocumentData
	if err := rows.Scan(&record.UploadID, &record.Path, &rawData); err != nil {
		return QualifiedDocumentData{}, err
	}

	data, err := s.serializer.UnmarshalDocumentData(rawData)
	if err != nil {
		return QualifiedDocumentData{}, err
	}
	record.Document = data

	if columns&DocumentDataColumnRanges == 0 {
		record.Document.Ranges = nil
	}
	if columns&DocumentDataColumnHoverText == 0 {
		record.Document.HoverResults = nil
	}
	if columns&DocumentDataColumnMonikers == 0 {
		record.Document.Monikers = nil
	}
	if columns&DocumentDataColumnPackageInformation == 0 {
		record.Document.PackageInformation = nil
	}
	if columns&DocumentDataColumnDiagnostics == 0 {
		record.Document.Diagnostics = nil
	}

	return record, nil
}

// makeResultChunkVisitor returns a function that accepts a mapping function, reads
// result chunk values from the given row object and calls the mapping function on
// each decoded result set.
func (s *Store) makeResultChunkVisitor(rows *sql.Rows, queryErr error) func(func(int, semantic.ResultChunkData)) error {
	return func(f func(int, semantic.ResultChunkData)) (err error) {
		if queryErr != nil {
			return queryErr
		}
		defer func() { err = basestore.CloseRows(rows, err) }()

		var rawData []byte
		for rows.Next() {
			var index int
			if err := rows.Scan(&index, &rawData); err != nil {
				return err
			}

			data, err := s.serializer.UnmarshalResultChunkData(rawData)
			if err != nil {
				return err
			}

			f(index, data)
		}

		return nil
	}
}

type QualifiedMonikerLocations struct {
	DumpID int
	semantic.MonikerLocations
}

// scanQualifiedMonikerLocations reads moniker locations values from the given row object.
func (s *Store) scanQualifiedMonikerLocations(rows *sql.Rows, queryErr error) (_ []QualifiedMonikerLocations, err error) {
	if queryErr != nil {
		return nil, queryErr
	}
	defer func() { err = basestore.CloseRows(rows, err) }()

	var values []QualifiedMonikerLocations
	for rows.Next() {
		record, err := s.scanSingleQualifiedMonikerLocationsObject(rows)
		if err != nil {
			return nil, err
		}

		values = append(values, record)
	}

	return values, nil
}

// scanSingleQualifiedMonikerLocationsObject populates a qualified moniker locations value
// from the given cursor.
func (s *Store) scanSingleQualifiedMonikerLocationsObject(rows *sql.Rows) (QualifiedMonikerLocations, error) {
	var rawData []byte
	var record QualifiedMonikerLocations
	if err := rows.Scan(&record.DumpID, &record.Scheme, &record.Identifier, &rawData); err != nil {
		return QualifiedMonikerLocations{}, err
	}

	data, err := s.serializer.UnmarshalLocations(rawData)
	if err != nil {
		return QualifiedMonikerLocations{}, err
	}
	record.Locations = data

	return record, nil
}
