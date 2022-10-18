package controllers

import (
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	sourcev1 "github.com/fluxcd/source-controller/api/v1beta2"
)

type SourceRevisionChangePredicate struct {
	predicate.Funcs
}

// Create implements Predicate.
func (SourceRevisionChangePredicate) Create(e event.CreateEvent) bool {
	sourceObj, ok := e.Object.(sourcev1.Source)
	if !ok {
		return false
	}

	if sourceObj.GetArtifact() != nil {
		return true
	}

	return false
}

// Delete implements Predicate.
func (SourceRevisionChangePredicate) Delete(e event.DeleteEvent) bool {
	return false
}

func (SourceRevisionChangePredicate) Update(e event.UpdateEvent) bool {
	if e.ObjectOld == nil || e.ObjectNew == nil {
		return false
	}

	oldSource, ok := e.ObjectOld.(sourcev1.Source)
	if !ok {
		return false
	}

	newSource, ok := e.ObjectNew.(sourcev1.Source)
	if !ok {
		return false
	}

	if oldSource.GetArtifact() == nil && newSource.GetArtifact() != nil {
		return true
	}

	if oldSource.GetArtifact() != nil && newSource.GetArtifact() != nil &&
		oldSource.GetArtifact().Revision != newSource.GetArtifact().Revision {
		return true
	}

	return false
}

// Generic implements Predicate.
func (SourceRevisionChangePredicate) Generic(e event.GenericEvent) bool {
	return false
}
