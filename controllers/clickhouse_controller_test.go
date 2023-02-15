// Copyright (c) 2022 Aiven, Helsinki, Finland. https://aiven.io/

package controllers

import (
	"context"
	"os"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/aiven/aiven-operator/api/v1alpha1"
	clickhouseuserconfig "github.com/aiven/aiven-operator/api/v1alpha1/userconfigs/clickhouse"
)

var _ = Describe("Clickhouse Controller", func() {
	// Define utility constants for object names and testing timeouts/durations and intervals.
	const (
		namespace = "default"

		timeout  = time.Minute * 20
		interval = time.Second * 10
	)

	var (
		os          *v1alpha1.Clickhouse
		serviceName string
		ctx         context.Context
	)

	BeforeEach(func() {
		ctx = context.Background()
		serviceName = "k8s-test-ch-acc-" + generateRandomID()
		os = chSpec(serviceName, namespace)

		By("Creating a new Clickhouse CR instance")
		Expect(k8sClient.Create(ctx, os)).Should(Succeed())

		rLookupKey := types.NamespacedName{Name: serviceName, Namespace: namespace}
		createdClickhouse := &v1alpha1.Clickhouse{}
		// We'll need to retry getting this newly created Clickhouse,
		// given that creation may not immediately happen.
		By("by retrieving os instance from k8s")
		Eventually(func() bool {
			err := k8sClient.Get(ctx, rLookupKey, createdClickhouse)

			return err == nil
		}, timeout, interval).Should(BeTrue())

		By("by waiting Clickhouse service status to become RUNNING")
		Eventually(func() bool {
			err := k8sClient.Get(ctx, rLookupKey, createdClickhouse)
			if err == nil {
				return meta.IsStatusConditionTrue(createdClickhouse.Status.Conditions, conditionTypeRunning)
			}
			return false
		}, timeout, interval).Should(BeTrue())

		By("by checking finalizers")
		Expect(createdClickhouse.GetFinalizers()).ToNot(BeEmpty())
	})

	Context("Validating Clickhouse reconciler behaviour", func() {
		It("should createOrUpdate a new Clickhouse service", func() {
			createdClickhouse := &v1alpha1.Clickhouse{}
			lookupKey := types.NamespacedName{Name: serviceName, Namespace: namespace}

			Expect(k8sClient.Get(ctx, lookupKey, createdClickhouse)).Should(Succeed())

			By("by checking that after creation of a Clickhouse service secret is created")
			createdSecret := &corev1.Secret{}
			Expect(k8sClient.Get(ctx, types.NamespacedName{Name: serviceName, Namespace: namespace}, createdSecret)).Should(Succeed())

			Expect(createdSecret.Data["HOST"]).NotTo(BeEmpty())
			Expect(createdSecret.Data["PORT"]).NotTo(BeEmpty())
			Expect(createdSecret.Data["USER"]).NotTo(BeEmpty())
			Expect(createdSecret.Data["PASSWORD"]).NotTo(BeEmpty())

			Expect(createdClickhouse.Status.State).Should(Equal("RUNNING"))

			// Userconfig test
			expectedIPFilter := []*clickhouseuserconfig.IpFilter{
				{
					Network: "10.20.0.0/16",
				},
				{
					Network:     "0.0.0.0",
					Description: anyPointer("whatever"),
				},
			}
			Expect(createdClickhouse.Spec.UserConfig.IpFilter).Should(Equal(expectedIPFilter))
		})
	})

	AfterEach(func() {
		By("Ensures that Clickhouse instance was deleted")
		ensureDelete(ctx, os)
	})
})

func chSpec(serviceName, namespace string) *v1alpha1.Clickhouse {
	return &v1alpha1.Clickhouse{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "aiven.io/v1alpha1",
			Kind:       "Clickhouse",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      serviceName,
			Namespace: namespace,
		},
		Spec: v1alpha1.ClickhouseSpec{
			ServiceCommonSpec: v1alpha1.ServiceCommonSpec{
				Project:   os.Getenv("AIVEN_PROJECT_NAME"),
				Plan:      "business-16",
				CloudName: "google-europe-west1",
				Tags:      map[string]string{"key1": "value1"},
			},
			UserConfig: &clickhouseuserconfig.ClickhouseUserConfig{
				IpFilter: []*clickhouseuserconfig.IpFilter{
					{
						Network: "10.20.0.0/16",
					},
					{
						Network:     "0.0.0.0",
						Description: anyPointer("whatever"),
					},
				},
			},
			AuthSecretRef: v1alpha1.AuthSecretReference{
				Name: secretRefName,
				Key:  secretRefKey,
			},
		},
	}
}