// Code generated by entc, DO NOT EDIT.

package item

import (
	"inventory-app/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Desc applies equality check predicate on the "desc" field. It's identical to DescEQ.
func Desc(v string) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDesc), v))
	})
}

// Count applies equality check predicate on the "count" field. It's identical to CountEQ.
func Count(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCount), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// DescEQ applies the EQ predicate on the "desc" field.
func DescEQ(v string) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDesc), v))
	})
}

// DescNEQ applies the NEQ predicate on the "desc" field.
func DescNEQ(v string) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDesc), v))
	})
}

// DescIn applies the In predicate on the "desc" field.
func DescIn(vs ...string) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDesc), v...))
	})
}

// DescNotIn applies the NotIn predicate on the "desc" field.
func DescNotIn(vs ...string) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDesc), v...))
	})
}

// DescGT applies the GT predicate on the "desc" field.
func DescGT(v string) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDesc), v))
	})
}

// DescGTE applies the GTE predicate on the "desc" field.
func DescGTE(v string) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDesc), v))
	})
}

// DescLT applies the LT predicate on the "desc" field.
func DescLT(v string) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDesc), v))
	})
}

// DescLTE applies the LTE predicate on the "desc" field.
func DescLTE(v string) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDesc), v))
	})
}

// DescContains applies the Contains predicate on the "desc" field.
func DescContains(v string) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDesc), v))
	})
}

// DescHasPrefix applies the HasPrefix predicate on the "desc" field.
func DescHasPrefix(v string) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDesc), v))
	})
}

// DescHasSuffix applies the HasSuffix predicate on the "desc" field.
func DescHasSuffix(v string) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDesc), v))
	})
}

// DescIsNil applies the IsNil predicate on the "desc" field.
func DescIsNil() predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDesc)))
	})
}

// DescNotNil applies the NotNil predicate on the "desc" field.
func DescNotNil() predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDesc)))
	})
}

// DescEqualFold applies the EqualFold predicate on the "desc" field.
func DescEqualFold(v string) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDesc), v))
	})
}

// DescContainsFold applies the ContainsFold predicate on the "desc" field.
func DescContainsFold(v string) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDesc), v))
	})
}

// TagsIsNil applies the IsNil predicate on the "tags" field.
func TagsIsNil() predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTags)))
	})
}

// TagsNotNil applies the NotNil predicate on the "tags" field.
func TagsNotNil() predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTags)))
	})
}

// CountEQ applies the EQ predicate on the "count" field.
func CountEQ(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCount), v))
	})
}

// CountNEQ applies the NEQ predicate on the "count" field.
func CountNEQ(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCount), v))
	})
}

// CountIn applies the In predicate on the "count" field.
func CountIn(vs ...int) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCount), v...))
	})
}

// CountNotIn applies the NotIn predicate on the "count" field.
func CountNotIn(vs ...int) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCount), v...))
	})
}

// CountGT applies the GT predicate on the "count" field.
func CountGT(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCount), v))
	})
}

// CountGTE applies the GTE predicate on the "count" field.
func CountGTE(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCount), v))
	})
}

// CountLT applies the LT predicate on the "count" field.
func CountLT(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCount), v))
	})
}

// CountLTE applies the LTE predicate on the "count" field.
func CountLTE(v int) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCount), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Item {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Item(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasGroup applies the HasEdge predicate on the "group" edge.
func HasGroup() predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, GroupTable, GroupPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupWith applies the HasEdge predicate on the "group" edge with a given conditions (other predicates).
func HasGroupWith(preds ...predicate.Group) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, GroupTable, GroupPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTag applies the HasEdge predicate on the "tag" edge.
func HasTag() predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TagTable, TagPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagWith applies the HasEdge predicate on the "tag" edge with a given conditions (other predicates).
func HasTagWith(preds ...predicate.Tag) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TagTable, TagPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Item) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Item) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
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
func Not(p predicate.Item) predicate.Item {
	return predicate.Item(func(s *sql.Selector) {
		p(s.Not())
	})
}
