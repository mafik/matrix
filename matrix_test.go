package matrix

import (
	"math"
	"testing"

	"mrogalski.eu/go/vec"
)

func eq(a, b vec.Vec) bool {
	return a.Sub(b).Len() < 0.001
}

func matrixDiffers(m, n Matrix) bool {
	for i := 0; i < 6; i++ {
		if m[i] != n[i] {
			return true
		}
	}
	return false
}

func TestIdentity(t *testing.T) {
	if matrixDiffers(Identity, Identity.Mul(Identity)) {
		t.Fail()
	}
}

func TestTranslation(t *testing.T) {
	m := Translation(vec.New(3, 5))
	if !m.Transform(vec.New(-1, -2)).Equal(vec.New(2, 3)) {
		t.Fail()
	}
}

func TestRotation(t *testing.T) {
	m := Rotation(math.Pi / 2)
	expected := vec.New(-2, 1)
	actual := m.Transform(vec.New(1, 2))
	if !eq(actual, expected) {
		t.Log("Expected", expected)
		t.Error("Found", actual)
	}
}

func TestDeterminant(t *testing.T) {
	m := Scale(2)
	if m.Determinant() != 4 {
		t.Fail()
	}
}

func TestInverse(t *testing.T) {
	m := Scale(2)
	if matrixDiffers(m.Inverse(), Scale(0.5)) {
		t.Error(m, Scale(0.5))
	}
	m = Translation(vec.New(2, 3))
	if matrixDiffers(m.Inverse(), Translation(vec.New(-2, -3))) {
		t.Error(m, Translation(vec.New(-2, -3)))
	}
	m = Rotation(0.5)
	if matrixDiffers(m.Inverse(), Rotation(-0.5)) {
		t.Error(m, Rotation(-0.5))
	}
}

func TestChainedTransforms(t *testing.T) {
	m := Identity
	test := func(description string, expected vec.Vec) {
		if !eq(m.Transform(vec.Zero), expected) {
			t.Log(description)
			t.Log("Expected", expected)
			t.Error("Found", m.Transform(vec.Zero))
		}
	}
	m = m.Mul(Translation(vec.New(1, 2)))
	test("Translation", vec.New(1, 2))
	m = m.Mul(Scale(3))
	test("Scale", vec.New(3, 6))
	m = m.Mul(Rotation(math.Pi / 2))
	test("Rotation", vec.New(-6, 3))
	m = Translation(vec.New(1, 0)).Mul(m)
	test("Pre-translation", vec.New(-6, 6))
}
