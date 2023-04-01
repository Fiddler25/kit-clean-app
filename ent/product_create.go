// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kit-clean-app/ent/order"
	"kit-clean-app/ent/product"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductCreate is the builder for creating a Product entity.
type ProductCreate struct {
	config
	mutation *ProductMutation
	hooks    []Hook
}

// SetCreated sets the "created" field.
func (pc *ProductCreate) SetCreated(t time.Time) *ProductCreate {
	pc.mutation.SetCreated(t)
	return pc
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (pc *ProductCreate) SetNillableCreated(t *time.Time) *ProductCreate {
	if t != nil {
		pc.SetCreated(*t)
	}
	return pc
}

// SetUpdated sets the "updated" field.
func (pc *ProductCreate) SetUpdated(t time.Time) *ProductCreate {
	pc.mutation.SetUpdated(t)
	return pc
}

// SetNillableUpdated sets the "updated" field if the given value is not nil.
func (pc *ProductCreate) SetNillableUpdated(t *time.Time) *ProductCreate {
	if t != nil {
		pc.SetUpdated(*t)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *ProductCreate) SetName(s string) *ProductCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *ProductCreate) SetDescription(s string) *ProductCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pc *ProductCreate) SetNillableDescription(s *string) *ProductCreate {
	if s != nil {
		pc.SetDescription(*s)
	}
	return pc
}

// SetPrice sets the "price" field.
func (pc *ProductCreate) SetPrice(f float64) *ProductCreate {
	pc.mutation.SetPrice(f)
	return pc
}

// SetStock sets the "stock" field.
func (pc *ProductCreate) SetStock(u uint8) *ProductCreate {
	pc.mutation.SetStock(u)
	return pc
}

// SetID sets the "id" field.
func (pc *ProductCreate) SetID(u uint32) *ProductCreate {
	pc.mutation.SetID(u)
	return pc
}

// AddOrderIDs adds the "orders" edge to the Order entity by IDs.
func (pc *ProductCreate) AddOrderIDs(ids ...uint32) *ProductCreate {
	pc.mutation.AddOrderIDs(ids...)
	return pc
}

// AddOrders adds the "orders" edges to the Order entity.
func (pc *ProductCreate) AddOrders(o ...*Order) *ProductCreate {
	ids := make([]uint32, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pc.AddOrderIDs(ids...)
}

// Mutation returns the ProductMutation object of the builder.
func (pc *ProductCreate) Mutation() *ProductMutation {
	return pc.mutation
}

// Save creates the Product in the database.
func (pc *ProductCreate) Save(ctx context.Context) (*Product, error) {
	pc.defaults()
	return withHooks[*Product, ProductMutation](ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProductCreate) SaveX(ctx context.Context) *Product {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProductCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProductCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProductCreate) defaults() {
	if _, ok := pc.mutation.Created(); !ok {
		v := product.DefaultCreated()
		pc.mutation.SetCreated(v)
	}
	if _, ok := pc.mutation.Updated(); !ok {
		v := product.DefaultUpdated()
		pc.mutation.SetUpdated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProductCreate) check() error {
	if _, ok := pc.mutation.Created(); !ok {
		return &ValidationError{Name: "created", err: errors.New(`ent: missing required field "Product.created"`)}
	}
	if _, ok := pc.mutation.Updated(); !ok {
		return &ValidationError{Name: "updated", err: errors.New(`ent: missing required field "Product.updated"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Product.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := product.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Product.name": %w`, err)}
		}
	}
	if v, ok := pc.mutation.Description(); ok {
		if err := product.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Product.description": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "Product.price"`)}
	}
	if v, ok := pc.mutation.Price(); ok {
		if err := product.PriceValidator(v); err != nil {
			return &ValidationError{Name: "price", err: fmt.Errorf(`ent: validator failed for field "Product.price": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Stock(); !ok {
		return &ValidationError{Name: "stock", err: errors.New(`ent: missing required field "Product.stock"`)}
	}
	if v, ok := pc.mutation.Stock(); ok {
		if err := product.StockValidator(v); err != nil {
			return &ValidationError{Name: "stock", err: fmt.Errorf(`ent: validator failed for field "Product.stock": %w`, err)}
		}
	}
	return nil
}

func (pc *ProductCreate) sqlSave(ctx context.Context) (*Product, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProductCreate) createSpec() (*Product, *sqlgraph.CreateSpec) {
	var (
		_node = &Product{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(product.Table, sqlgraph.NewFieldSpec(product.FieldID, field.TypeUint32))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Created(); ok {
		_spec.SetField(product.FieldCreated, field.TypeTime, value)
		_node.Created = value
	}
	if value, ok := pc.mutation.Updated(); ok {
		_spec.SetField(product.FieldUpdated, field.TypeTime, value)
		_node.Updated = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(product.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(product.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.Price(); ok {
		_spec.SetField(product.FieldPrice, field.TypeFloat64, value)
		_node.Price = value
	}
	if value, ok := pc.mutation.Stock(); ok {
		_spec.SetField(product.FieldStock, field.TypeUint8, value)
		_node.Stock = value
	}
	if nodes := pc.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.OrdersTable,
			Columns: []string{product.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProductCreateBulk is the builder for creating many Product entities in bulk.
type ProductCreateBulk struct {
	config
	builders []*ProductCreate
}

// Save creates the Product entities in the database.
func (pcb *ProductCreateBulk) Save(ctx context.Context) ([]*Product, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Product, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProductCreateBulk) SaveX(ctx context.Context) []*Product {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProductCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProductCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
