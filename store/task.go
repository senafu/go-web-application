package store

import (
	"context"

	"github.com/funasedaisuke/go-web-application/entity"
)

func (r *Repository) ListTasks(
	ctx context.Context, db Queryer,

)(entity.Tasks,error){
	tasks := entity.Tasks{}
	sql := `select id ,title,status,created, modified from task;`
	if err := db.SelectContext(ctx,&tasks,sql);err != nil{
		return nil, err
	}
	return tasks,nil
}


func (r *Repository) AddTasks(
	ctx context.Context, db Execer,t *entity.Task,

)error{
	t.Created = r.Clocker.Now()
	t.Modified = r.Clocker.Now()
	sql := `INSERT INTO task
	(title,status,created,modified)
	VALUES(?,?,?,?)`
	result,err := db.ExecContext(
		ctx,sql, t.Title,t.Status,
		t.Created,t.Modified,
	)
	if err != nil {
		return err
	}
	id ,err := result.LastInsertId()
	if err != nil{
		return err
	}
	t.ID = entity.TaskID(id)
	return nil
	
}