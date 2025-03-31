// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"bytes"
	"context"
	"errors"
	"sync/atomic"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/dlukt/graphql-backend-starter/ent"
	"github.com/rs/xid"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		schema:     cfg.Schema,
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Schema     *ast.Schema
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Mutation() MutationResolver
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Mutation struct {
		CreateProfile func(childComplexity int, input ent.CreateProfileInput) int
		DeleteProfile func(childComplexity int, id xid.ID) int
		UpdateProfile func(childComplexity int, id xid.ID, input ent.UpdateProfileInput) int
	}

	PageInfo struct {
		EndCursor       func(childComplexity int) int
		HasNextPage     func(childComplexity int) int
		HasPreviousPage func(childComplexity int) int
		StartCursor     func(childComplexity int) int
	}

	Profile struct {
		CreateTime func(childComplexity int) int
		Gender     func(childComplexity int) int
		ID         func(childComplexity int) int
		Name       func(childComplexity int) int
		Sub        func(childComplexity int) int
		UpdateTime func(childComplexity int) int
	}

	ProfileConnection struct {
		Edges      func(childComplexity int) int
		PageInfo   func(childComplexity int) int
		TotalCount func(childComplexity int) int
	}

	ProfileEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}

	Query struct {
		Node     func(childComplexity int, id xid.ID) int
		Nodes    func(childComplexity int, ids []xid.ID) int
		Profiles func(childComplexity int, after *entgql.Cursor[xid.ID], first *int, before *entgql.Cursor[xid.ID], last *int, orderBy *ent.ProfileOrder, where *ent.ProfileWhereInput) int
	}
}

type executableSchema struct {
	schema     *ast.Schema
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	if e.schema != nil {
		return e.schema
	}
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]any) (int, bool) {
	ec := executionContext{nil, e, 0, 0, nil}
	_ = ec
	switch typeName + "." + field {

	case "Mutation.createProfile":
		if e.complexity.Mutation.CreateProfile == nil {
			break
		}

		args, err := ec.field_Mutation_createProfile_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateProfile(childComplexity, args["input"].(ent.CreateProfileInput)), true

	case "Mutation.deleteProfile":
		if e.complexity.Mutation.DeleteProfile == nil {
			break
		}

		args, err := ec.field_Mutation_deleteProfile_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.DeleteProfile(childComplexity, args["id"].(xid.ID)), true

	case "Mutation.updateProfile":
		if e.complexity.Mutation.UpdateProfile == nil {
			break
		}

		args, err := ec.field_Mutation_updateProfile_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.UpdateProfile(childComplexity, args["id"].(xid.ID), args["input"].(ent.UpdateProfileInput)), true

	case "PageInfo.endCursor":
		if e.complexity.PageInfo.EndCursor == nil {
			break
		}

		return e.complexity.PageInfo.EndCursor(childComplexity), true

	case "PageInfo.hasNextPage":
		if e.complexity.PageInfo.HasNextPage == nil {
			break
		}

		return e.complexity.PageInfo.HasNextPage(childComplexity), true

	case "PageInfo.hasPreviousPage":
		if e.complexity.PageInfo.HasPreviousPage == nil {
			break
		}

		return e.complexity.PageInfo.HasPreviousPage(childComplexity), true

	case "PageInfo.startCursor":
		if e.complexity.PageInfo.StartCursor == nil {
			break
		}

		return e.complexity.PageInfo.StartCursor(childComplexity), true

	case "Profile.createTime":
		if e.complexity.Profile.CreateTime == nil {
			break
		}

		return e.complexity.Profile.CreateTime(childComplexity), true

	case "Profile.gender":
		if e.complexity.Profile.Gender == nil {
			break
		}

		return e.complexity.Profile.Gender(childComplexity), true

	case "Profile.id":
		if e.complexity.Profile.ID == nil {
			break
		}

		return e.complexity.Profile.ID(childComplexity), true

	case "Profile.name":
		if e.complexity.Profile.Name == nil {
			break
		}

		return e.complexity.Profile.Name(childComplexity), true

	case "Profile.sub":
		if e.complexity.Profile.Sub == nil {
			break
		}

		return e.complexity.Profile.Sub(childComplexity), true

	case "Profile.updateTime":
		if e.complexity.Profile.UpdateTime == nil {
			break
		}

		return e.complexity.Profile.UpdateTime(childComplexity), true

	case "ProfileConnection.edges":
		if e.complexity.ProfileConnection.Edges == nil {
			break
		}

		return e.complexity.ProfileConnection.Edges(childComplexity), true

	case "ProfileConnection.pageInfo":
		if e.complexity.ProfileConnection.PageInfo == nil {
			break
		}

		return e.complexity.ProfileConnection.PageInfo(childComplexity), true

	case "ProfileConnection.totalCount":
		if e.complexity.ProfileConnection.TotalCount == nil {
			break
		}

		return e.complexity.ProfileConnection.TotalCount(childComplexity), true

	case "ProfileEdge.cursor":
		if e.complexity.ProfileEdge.Cursor == nil {
			break
		}

		return e.complexity.ProfileEdge.Cursor(childComplexity), true

	case "ProfileEdge.node":
		if e.complexity.ProfileEdge.Node == nil {
			break
		}

		return e.complexity.ProfileEdge.Node(childComplexity), true

	case "Query.node":
		if e.complexity.Query.Node == nil {
			break
		}

		args, err := ec.field_Query_node_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Node(childComplexity, args["id"].(xid.ID)), true

	case "Query.nodes":
		if e.complexity.Query.Nodes == nil {
			break
		}

		args, err := ec.field_Query_nodes_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Nodes(childComplexity, args["ids"].([]xid.ID)), true

	case "Query.profiles":
		if e.complexity.Query.Profiles == nil {
			break
		}

		args, err := ec.field_Query_profiles_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Profiles(childComplexity, args["after"].(*entgql.Cursor[xid.ID]), args["first"].(*int), args["before"].(*entgql.Cursor[xid.ID]), args["last"].(*int), args["orderBy"].(*ent.ProfileOrder), args["where"].(*ent.ProfileWhereInput)), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	opCtx := graphql.GetOperationContext(ctx)
	ec := executionContext{opCtx, e, 0, 0, make(chan graphql.DeferredResult)}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputCreateProfileInput,
		ec.unmarshalInputProfileOrder,
		ec.unmarshalInputProfileWhereInput,
		ec.unmarshalInputUpdateProfileInput,
	)
	first := true

	switch opCtx.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			var response graphql.Response
			var data graphql.Marshaler
			if first {
				first = false
				ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
				data = ec._Query(ctx, opCtx.Operation.SelectionSet)
			} else {
				if atomic.LoadInt32(&ec.pendingDeferred) > 0 {
					result := <-ec.deferredResults
					atomic.AddInt32(&ec.pendingDeferred, -1)
					data = result.Result
					response.Path = result.Path
					response.Label = result.Label
					response.Errors = result.Errors
				} else {
					return nil
				}
			}
			var buf bytes.Buffer
			data.MarshalGQL(&buf)
			response.Data = buf.Bytes()
			if atomic.LoadInt32(&ec.deferred) > 0 {
				hasNext := atomic.LoadInt32(&ec.pendingDeferred) > 0
				response.HasNext = &hasNext
			}

			return &response
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Mutation(ctx, opCtx.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
	deferred        int32
	pendingDeferred int32
	deferredResults chan graphql.DeferredResult
}

func (ec *executionContext) processDeferredGroup(dg graphql.DeferredGroup) {
	atomic.AddInt32(&ec.pendingDeferred, 1)
	go func() {
		ctx := graphql.WithFreshResponseContext(dg.Context)
		dg.FieldSet.Dispatch(ctx)
		ds := graphql.DeferredResult{
			Path:   dg.Path,
			Label:  dg.Label,
			Result: dg.FieldSet,
			Errors: graphql.GetErrors(ctx),
		}
		// null fields should bubble up
		if dg.FieldSet.Invalids > 0 {
			ds.Result = graphql.Null
		}
		ec.deferredResults <- ds
	}()
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(ec.Schema()), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(ec.Schema(), ec.Schema().Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "../../ent.graphql", Input: `directive @goField(forceResolver: Boolean, name: String, omittable: Boolean) on FIELD_DEFINITION | INPUT_FIELD_DEFINITION
directive @goModel(model: String, models: [String!], forceGenerate: Boolean) on OBJECT | INPUT_OBJECT | SCALAR | ENUM | INTERFACE | UNION
"""
CreateProfileInput is used for create Profile object.
Input was generated by ent.
"""
input CreateProfileInput {
  createTime: Time
  updateTime: Time
  sub: String!
  name: String!
  gender: String!
}
"""
Define a Relay Cursor type:
https://relay.dev/graphql/connections.htm#sec-Cursor
"""
scalar Cursor
"""
An object with an ID.
Follows the [Relay Global Object Identification Specification](https://relay.dev/graphql/objectidentification.htm)
"""
interface Node @goModel(model: "github.com/dlukt/graphql-backend-starter/ent.Noder") {
  """
  The id of the object.
  """
  id: ID!
}
"""
Possible directions in which to order a list of items when provided an ` + "`" + `orderBy` + "`" + ` argument.
"""
enum OrderDirection {
  """
  Specifies an ascending order for a given ` + "`" + `orderBy` + "`" + ` argument.
  """
  ASC
  """
  Specifies a descending order for a given ` + "`" + `orderBy` + "`" + ` argument.
  """
  DESC
}
"""
Information about pagination in a connection.
https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo
"""
type PageInfo {
  """
  When paginating forwards, are there more items?
  """
  hasNextPage: Boolean!
  """
  When paginating backwards, are there more items?
  """
  hasPreviousPage: Boolean!
  """
  When paginating backwards, the cursor to continue.
  """
  startCursor: Cursor
  """
  When paginating forwards, the cursor to continue.
  """
  endCursor: Cursor
}
type Profile implements Node {
  id: ID!
  createTime: Time!
  updateTime: Time
  sub: String!
  name: String!
  gender: String!
}
"""
A connection to a list of items.
"""
type ProfileConnection {
  """
  A list of edges.
  """
  edges: [ProfileEdge]
  """
  Information to aid in pagination.
  """
  pageInfo: PageInfo!
  """
  Identifies the total count of items in the connection.
  """
  totalCount: Int!
}
"""
An edge in a connection.
"""
type ProfileEdge {
  """
  The item at the end of the edge.
  """
  node: Profile
  """
  A cursor for use in pagination.
  """
  cursor: Cursor!
}
"""
Ordering options for Profile connections
"""
input ProfileOrder {
  """
  The ordering direction.
  """
  direction: OrderDirection! = ASC
  """
  The field by which to order Profiles.
  """
  field: ProfileOrderField!
}
"""
Properties by which Profile connections can be ordered.
"""
enum ProfileOrderField {
  CREATE_TIME
  UPDATE_TIME
  NAME
}
"""
ProfileWhereInput is used for filtering Profile objects.
Input was generated by ent.
"""
input ProfileWhereInput {
  not: ProfileWhereInput
  and: [ProfileWhereInput!]
  or: [ProfileWhereInput!]
  """
  id field predicates
  """
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """
  create_time field predicates
  """
  createTime: Time
  createTimeNEQ: Time
  createTimeIn: [Time!]
  createTimeNotIn: [Time!]
  createTimeGT: Time
  createTimeGTE: Time
  createTimeLT: Time
  createTimeLTE: Time
  """
  update_time field predicates
  """
  updateTime: Time
  updateTimeNEQ: Time
  updateTimeIn: [Time!]
  updateTimeNotIn: [Time!]
  updateTimeGT: Time
  updateTimeGTE: Time
  updateTimeLT: Time
  updateTimeLTE: Time
  updateTimeIsNil: Boolean
  updateTimeNotNil: Boolean
  """
  sub field predicates
  """
  sub: String
  subNEQ: String
  subIn: [String!]
  subNotIn: [String!]
  subGT: String
  subGTE: String
  subLT: String
  subLTE: String
  subContains: String
  subHasPrefix: String
  subHasSuffix: String
  subEqualFold: String
  subContainsFold: String
  """
  name field predicates
  """
  name: String
  nameNEQ: String
  nameIn: [String!]
  nameNotIn: [String!]
  nameGT: String
  nameGTE: String
  nameLT: String
  nameLTE: String
  nameContains: String
  nameHasPrefix: String
  nameHasSuffix: String
  nameEqualFold: String
  nameContainsFold: String
  """
  gender field predicates
  """
  gender: String
  genderNEQ: String
  genderIn: [String!]
  genderNotIn: [String!]
  genderGT: String
  genderGTE: String
  genderLT: String
  genderLTE: String
  genderContains: String
  genderHasPrefix: String
  genderHasSuffix: String
  genderEqualFold: String
  genderContainsFold: String
}
type Query {
  """
  Fetches an object given its ID.
  """
  node(
    """
    ID of the object.
    """
    id: ID!
  ): Node
  """
  Lookup nodes by a list of IDs.
  """
  nodes(
    """
    The list of node IDs.
    """
    ids: [ID!]!
  ): [Node]!
  profiles(
    """
    Returns the elements in the list that come after the specified cursor.
    """
    after: Cursor

    """
    Returns the first _n_ elements from the list.
    """
    first: Int

    """
    Returns the elements in the list that come before the specified cursor.
    """
    before: Cursor

    """
    Returns the last _n_ elements from the list.
    """
    last: Int

    """
    Ordering options for Profiles returned from the connection.
    """
    orderBy: ProfileOrder

    """
    Filtering options for Profiles returned from the connection.
    """
    where: ProfileWhereInput
  ): ProfileConnection!
}
"""
The builtin Time type
"""
scalar Time
"""
UpdateProfileInput is used for update Profile object.
Input was generated by ent.
"""
input UpdateProfileInput {
  updateTime: Time
  clearUpdateTime: Boolean
  sub: String
  name: String
  gender: String
}
`, BuiltIn: false},
	{Name: "../../starter.graphql", Input: `type Mutation {
  createProfile(input: CreateProfileInput!): Profile
  updateProfile(id: ID!, input: UpdateProfileInput!): Profile
  deleteProfile(id: ID!): ID
}
`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
