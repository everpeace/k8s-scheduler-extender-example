package main

import (
	"fmt"
	"github.com/wangchen615/k8s-scheduler-extender-example/k8s-scheduler-extender"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/comail/colog"
	"github.com/julienschmidt/httprouter"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
)

var (
	TruePredicate = k8s_scheduler_extender.Predicate{
		Name: "always_true",
		Func: func(pod v1.Pod, node v1.Node) (bool, error) {
			return true, nil
		},
	}

	ZeroPriority = k8s_scheduler_extender.Prioritize{
		Name: "zero_score",
		Func: func(_ v1.Pod, nodes []v1.Node) (*schedulerapi.HostPriorityList, error) {
			var priorityList schedulerapi.HostPriorityList
			priorityList = make([]schedulerapi.HostPriority, len(nodes))
			for i, node := range nodes {
				priorityList[i] = schedulerapi.HostPriority{
					Host:  node.Name,
					Score: 0,
				}
			}
			return &priorityList, nil
		},
	}

	NoBind = k8s_scheduler_extender.Bind{
		Func: func(podName string, podNamespace string, podUID types.UID, node string) error {
			return fmt.Errorf("This extender doesn't support Bind.  Please make 'BindVerb' be empty in your ExtenderConfig.")
		},
	}

	EchoPreemption = k8s_scheduler_extender.Preemption{
		Func: func(
			_ v1.Pod,
			_ map[string]*schedulerapi.Victims,
			nodeNameToMetaVictims map[string]*schedulerapi.MetaVictims,
		) map[string]*schedulerapi.MetaVictims {
			return nodeNameToMetaVictims
		},
	}
)

func main() {
	colog.SetDefaultLevel(colog.LInfo)
	colog.SetMinLevel(colog.LInfo)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()
	level := k8s_scheduler_extender.StringToLevel(os.Getenv("LOG_LEVEL"))
	log.Print("Log level was set to ", strings.ToUpper(level.String()))
	colog.SetMinLevel(level)

	router := httprouter.New()
	k8s_scheduler_extender.AddVersion(router)

	predicates := []k8s_scheduler_extender.Predicate{TruePredicate}
	for _, p := range predicates {
		k8s_scheduler_extender.AddPredicate(router, p)
	}

	priorities := []k8s_scheduler_extender.Prioritize{ZeroPriority}
	for _, p := range priorities {
		k8s_scheduler_extender.AddPrioritize(router, p)
	}

	k8s_scheduler_extender.AddBind(router, NoBind)

	log.Print("info: server starting on the port :80")
	if err := http.ListenAndServe(":80", router); err != nil {
		log.Fatal(err)
	}
}
