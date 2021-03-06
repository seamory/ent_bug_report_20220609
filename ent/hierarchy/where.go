// Code generated by entc, DO NOT EDIT.

package hierarchy

import (
	"entgo.io/bug/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Describe applies equality check predicate on the "describe" field. It's identical to DescribeEQ.
func Describe(v string) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescribe), v))
	})
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v int) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSort), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Hierarchy {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hierarchy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Hierarchy {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hierarchy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// DescribeEQ applies the EQ predicate on the "describe" field.
func DescribeEQ(v string) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescribe), v))
	})
}

// DescribeNEQ applies the NEQ predicate on the "describe" field.
func DescribeNEQ(v string) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescribe), v))
	})
}

// DescribeIn applies the In predicate on the "describe" field.
func DescribeIn(vs ...string) predicate.Hierarchy {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hierarchy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescribe), v...))
	})
}

// DescribeNotIn applies the NotIn predicate on the "describe" field.
func DescribeNotIn(vs ...string) predicate.Hierarchy {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hierarchy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescribe), v...))
	})
}

// DescribeGT applies the GT predicate on the "describe" field.
func DescribeGT(v string) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescribe), v))
	})
}

// DescribeGTE applies the GTE predicate on the "describe" field.
func DescribeGTE(v string) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescribe), v))
	})
}

// DescribeLT applies the LT predicate on the "describe" field.
func DescribeLT(v string) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescribe), v))
	})
}

// DescribeLTE applies the LTE predicate on the "describe" field.
func DescribeLTE(v string) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescribe), v))
	})
}

// DescribeContains applies the Contains predicate on the "describe" field.
func DescribeContains(v string) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescribe), v))
	})
}

// DescribeHasPrefix applies the HasPrefix predicate on the "describe" field.
func DescribeHasPrefix(v string) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescribe), v))
	})
}

// DescribeHasSuffix applies the HasSuffix predicate on the "describe" field.
func DescribeHasSuffix(v string) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescribe), v))
	})
}

// DescribeEqualFold applies the EqualFold predicate on the "describe" field.
func DescribeEqualFold(v string) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescribe), v))
	})
}

// DescribeContainsFold applies the ContainsFold predicate on the "describe" field.
func DescribeContainsFold(v string) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescribe), v))
	})
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v int) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSort), v))
	})
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v int) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSort), v))
	})
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...int) predicate.Hierarchy {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hierarchy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSort), v...))
	})
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...int) predicate.Hierarchy {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hierarchy(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSort), v...))
	})
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v int) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSort), v))
	})
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v int) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSort), v))
	})
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v int) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSort), v))
	})
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v int) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSort), v))
	})
}

// HasChild applies the HasEdge predicate on the "child" edge.
func HasChild() predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChildTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ChildTable, ChildPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildWith applies the HasEdge predicate on the "child" edge with a given conditions (other predicates).
func HasChildWith(preds ...predicate.Hierarchy) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ChildTable, ChildPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ParentTable, ParentPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Hierarchy) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ParentTable, ParentPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Hierarchy) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Hierarchy) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Hierarchy) predicate.Hierarchy {
	return predicate.Hierarchy(func(s *sql.Selector) {
		p(s.Not())
	})
}
