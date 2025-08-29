package rules_test

import (
    "context"
    "testing"

    "entgo.io/ent/dialect"
    "github.com/dlukt/graphql-backend-starter/ent"
    _ "github.com/dlukt/graphql-backend-starter/ent/runtime"
    "github.com/dlukt/graphql-backend-starter/ent/profile"
    "github.com/dlukt/graphql-backend-starter/rules/claims"
    "github.com/dlukt/graphql-backend-starter/rules/viewer"
    _ "github.com/mattn/go-sqlite3"
)

func newTestClient(t *testing.T) *ent.Client {
    t.Helper()
    client, err := ent.Open(dialect.SQLite, "file:ent.db?mode=memory&cache=shared&_fk=1")
    if err != nil {
        t.Fatalf("failed opening connection to sqlite: %v", err)
    }
    t.Cleanup(func() { _ = client.Close() })
    if err := client.Schema.Create(context.Background()); err != nil {
        t.Fatalf("failed creating schema resources: %v", err)
    }
    return client
}

func viewerCtx(ctx context.Context, sub string) context.Context {
    c := &claims.Claims{Sub: sub}
    v := viewer.NewFromClaims(c)
    return viewer.NewContext(ctx, v)
}

func TestProfileCreate_AllowsAuthenticated(t *testing.T) {
    client := newTestClient(t)
    base := context.Background()
    ctx := ent.NewContext(viewerCtx(base, "user-1"), client)

    p, err := client.Profile.Create().Save(ctx)
    if err != nil {
        t.Fatalf("expected create to succeed, got error: %v", err)
    }
    if p.Sub != "user-1" {
        t.Fatalf("expected sub to be set from viewer, got %q", p.Sub)
    }
}

func TestProfileUpdateOne_AllowsOwner(t *testing.T) {
    client := newTestClient(t)
    base := context.Background()
    ctx := ent.NewContext(viewerCtx(base, "user-2"), client)

    // Seed a profile for this owner
    p := client.Profile.Create().SaveX(ctx)

    // Owner updates their own profile
    if _, err := client.Profile.UpdateOneID(p.ID).SetName("Updated").Save(ctx); err != nil {
        t.Fatalf("expected update to succeed for owner, got error: %v", err)
    }

    // Verify updated
    got := client.Profile.GetX(ctx, p.ID)
    if got.Name != "Updated" {
        t.Fatalf("expected name to be Updated, got: %v", got.Name)
    }
}

func TestProfileDeleteOne_AllowsOwner(t *testing.T) {
    client := newTestClient(t)
    base := context.Background()
    ctx := ent.NewContext(viewerCtx(base, "user-3"), client)

    // Seed a profile for this owner
    p := client.Profile.Create().SaveX(ctx)

    // Owner deletes their own profile
    if err := client.Profile.DeleteOneID(p.ID).Exec(ctx); err != nil {
        t.Fatalf("expected delete to succeed for owner, got error: %v", err)
    }

    // Verify deleted
    if n := client.Profile.Query().Where(profile.ID(p.ID)).CountX(ctx); n != 0 {
        t.Fatalf("expected 0 profiles after delete, got %d", n)
    }
}

func TestProfileCreate_DeniesUnauthenticated(t *testing.T) {
    client := newTestClient(t)
    // No viewer/claims on context
    ctx := ent.NewContext(context.Background(), client)

    if _, err := client.Profile.Create().Save(ctx); err == nil {
        t.Fatalf("expected unauthenticated create to be denied")
    }
}

func TestProfileUpdateOne_DeniesNonOwner(t *testing.T) {
    client := newTestClient(t)
    base := context.Background()
    ownerCtx := ent.NewContext(viewerCtx(base, "owner-1"), client)

    // Seed as owner
    p := client.Profile.Create().SaveX(ownerCtx)

    // Intruder tries to update
    intruderCtx := ent.NewContext(viewerCtx(base, "intruder-1"), client)
    if _, err := client.Profile.UpdateOneID(p.ID).SetName("Hacked").Save(intruderCtx); err == nil {
        t.Fatalf("expected non-owner update to be denied")
    }
    // Verify unchanged
    got := client.Profile.GetX(ownerCtx, p.ID)
    if got.Name == "Hacked" {
        t.Fatalf("name was unexpectedly changed by non-owner")
    }
}

func TestProfileDeleteOne_DeniesNonOwner(t *testing.T) {
    client := newTestClient(t)
    base := context.Background()
    ownerCtx := ent.NewContext(viewerCtx(base, "owner-2"), client)

    // Seed as owner
    p := client.Profile.Create().SaveX(ownerCtx)

    // Intruder tries to delete
    intruderCtx := ent.NewContext(viewerCtx(base, "intruder-2"), client)
    if err := client.Profile.DeleteOneID(p.ID).Exec(intruderCtx); err == nil {
        t.Fatalf("expected non-owner delete to be denied")
    }
    // Verify still exists
    if n := client.Profile.Query().Where(profile.ID(p.ID)).CountX(ownerCtx); n != 1 {
        t.Fatalf("expected profile to remain after denied delete, count=%d", n)
    }
}

func TestProfileBulkUpdate_Denied(t *testing.T) {
    client := newTestClient(t)
    base := context.Background()
    ctx := ent.NewContext(viewerCtx(base, "bulk-user"), client)
    // Seed one profile
    _ = client.Profile.Create().SaveX(ctx)

    if _, err := client.Profile.Update().SetName("Bulk").Save(ctx); err == nil {
        t.Fatalf("expected bulk update to be denied")
    }
}

func TestProfileBulkDelete_Denied(t *testing.T) {
    client := newTestClient(t)
    base := context.Background()
    ctx := ent.NewContext(viewerCtx(base, "bulk-user-2"), client)
    // Seed one profile
    _ = client.Profile.Create().SaveX(ctx)

    if _, err := client.Profile.Delete().Exec(ctx); err == nil {
        t.Fatalf("expected bulk delete to be denied")
    }
}
